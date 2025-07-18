package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kryast/Crud-6.git/models"
	"github.com/kryast/Crud-6.git/services"
)

type CustomerHandler struct {
	svc services.CustomerService
}

func NewCustomerHandler(s services.CustomerService) *CustomerHandler {
	return &CustomerHandler{s}
}

func (h *CustomerHandler) Create(c *gin.Context) {
	var in models.Customer
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
