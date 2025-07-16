package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderID   uint    `json:"order_id" binding:"required"`
	Order     Order   `json:"order" gorm:"foreignKey:OrderID"`
	ProductID uint    `json:"product_id" binding:"required"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
	Quantity  int     `json:"quantity" binding:"required,min=1"`
	Subtotal  float64 `json:"subtotal"`
}
