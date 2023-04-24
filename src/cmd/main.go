package main

import (
	"context"
	"log"
	"net/http"
	"taynguyen/sample/src/cmd/router"
	"taynguyen/sample/src/handler"
	"taynguyen/sample/src/repo"
	"taynguyen/sample/src/utils"

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
	StartServer(h) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func StartServer(h *handler.Handler) *http.Server {
	r := router.CreateRouter(h)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	return srv
}
