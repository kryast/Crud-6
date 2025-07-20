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
	r.PUT("/customers/:id", h.Customer.Update)
	r.DELETE("/customers/:id", h.Customer.Delete)

	// Products
	r.POST("/products", h.Product.Create)
	r.GET("/products", h.Product.GetAll)
	r.GET("/products/:id", h.Product.GetByID)
	r.PUT("/products/:id", h.Product.Update)
	r.DELETE("/products/:id", h.Product.Delete)

	// Orders
	r.POST("/orders", h.Order.Create)
	r.GET("/orders", h.Order.GetAll)
	r.GET("/orders/:id", h.Order.GetByID)
	r.PUT("/orders/:id", h.Order.Update)
	r.DELETE("/orders/:id", h.Order.Delete)

	return r
}
