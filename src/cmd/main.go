package main

import (
	"context"
	"net/http"
	"taynguyen/sample/src/handler"
	"taynguyen/sample/src/repo"
	"taynguyen/sample/src/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

func main() {
	ctx := context.Background()

	// Initialize snowflake
	if err := utils.InitializeSnowFlake(); err != nil {
		panic(err)
	}

	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5435",
		User:     "taylor",
		Password: "taylor",
		Database: "taylor",
	})
	if err := db.Ping(ctx); err != nil {
		panic(err)
	}

	repo := repo.New(db)
	h := handler.New(repo)
	startServer(h) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func startServer(h *handler.Handler) {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/order", h.CreateOrder)
	r.GET("/orders", h.GetOrders)

	r.Run()
}
