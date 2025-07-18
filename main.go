package main

import (
	"log"

	"github.com/kryast/Crud-6.git/database"
	"github.com/kryast/Crud-6.git/handlers"
	"github.com/kryast/Crud-6.git/models"
	"github.com/kryast/Crud-6.git/repositories"
	"github.com/kryast/Crud-6.git/router"
	"github.com/kryast/Crud-6.git/services"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("DB connect fail: %v", err)
	}

	if err := db.AutoMigrate(&models.Customer{}, &models.Product{}, &models.Order{}, &models.OrderItem{}, &models.Payment{}); err != nil {
		log.Fatalf("migrate fail: %v", err)
	}

	custRepo := repositories.NewCustomerRepository(db)
	prodRepo := repositories.NewProductRepository(db)
	ordRepo := repositories.NewOrderRepository(db)
	oiRepo := repositories.NewOrderItemRepository(db)
	payRepo := repositories.NewPaymentRepository(db)

	custSvc := services.NewCustomerService(custRepo)
	prodSvc := services.NewProductService(prodRepo)
	ordSvc := services.NewOrderService(ordRepo)
	oiSvc := services.NewOrderItemService(oiRepo, prodRepo)
	paySvc := services.NewPaymentService(payRepo, oiRepo)

	custH := handlers.NewCustomerHandler(custSvc)
	prodH := handlers.NewProductHandler(prodSvc)
	ordH := handlers.NewOrderHandler(ordSvc)
	oiH := handlers.NewOrderItemHandler(oiSvc)
	payH := handlers.NewPaymentHandler(paySvc)

	r := router.SetupRouter(router.Handlers{
		Customer:  custH,
		Product:   prodH,
		Order:     ordH,
		OrderItem: oiH,
		Payment:   payH,
	})

	log.Println("🚀 API running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
