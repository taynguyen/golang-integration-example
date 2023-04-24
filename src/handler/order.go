package handler

import (
	"net/http"
	"taynguyen/sample/src/repo/orm"

	"github.com/gin-gonic/gin"
)

type NewOrderRq struct {
	Title      string `json:"title"`
	CustomerID int    `json:"customer_id"`
	Price      int    `json:"price"`
}

type OrderResponse struct {
	ID         int64  `json:"id"`
	Title      string `json:"title"`
	CustomerID int    `json:"customer_id"`
	Price      int    `json:"price"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

func (h Handler) CreateOrder(c *gin.Context) {
	o := &NewOrderRq{}

	if err := c.BindJSON(o); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}

	instance, err := h.repo.Order().Create(orm.Order{
		Title:      o.Title,
		CustomerId: o.CustomerID,
		Price:      o.Price,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
			"detail":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, orderToResponse(instance))
}

func (h Handler) GetOrders(c *gin.Context) {
	orders, err := h.repo.Order().GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
			"detail":  err.Error(),
		})
	}

	resModels := make([]OrderResponse, 0, len(orders))
	for _, o := range orders {
		resModels = append(resModels, orderToResponse(o))
	}

	c.JSON(http.StatusOK, resModels)
}

func orderToResponse(o orm.Order) OrderResponse {
	rs := OrderResponse{
		ID:         o.Id,
		Title:      o.Title,
		CustomerID: o.CustomerId,
		Price:      o.Price,
		CreatedAt:  o.CreatedAt.String(),
		UpdatedAt:  o.UpdatedAt.String(),
	}

	return rs
}
