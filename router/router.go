package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kryast/Crud-6.git/handlers"
)

type Handlers struct {
	Customer  *handlers.CustomerHandler
	Product   *handlers.ProductHandler
	Order     *handlers.OrderHandler
	OrderItem *handlers.OrderItemHandler
	Payment   *handlers.PaymentHandler
}

func SetupRouter(h Handlers) *gin.Engine {
	r := gin.Default()

	// Customers
	r.POST("/customers", h.Customer.Create)
	r.GET("/customers", h.Customer.GetAll)
	r.GET("/customers/:id", h.Customer.GetByID)

	return r
}
