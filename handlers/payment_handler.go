package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kryast/Crud-6.git/models"
	"github.com/kryast/Crud-6.git/services"
)

type PaymentHandler struct {
	svc services.PaymentService
}

func NewPaymentHandler(s services.PaymentService) *PaymentHandler {
	return &PaymentHandler{s}
}

func (h *PaymentHandler) Create(c *gin.Context) {
	var in models.Payment
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.svc.Create(&in); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, in)
}
