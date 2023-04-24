package router

import (
	"net/http"
	"taynguyen/sample/src/handler"

	"github.com/gin-gonic/gin"
)

func CreateRouter(h *handler.Handler) http.Handler {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/order", h.CreateOrder)
	r.GET("/orders", h.GetOrders)

	return r
}
