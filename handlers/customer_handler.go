package handlers

import (
	"net/http"
	"strconv"

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

func (h *CustomerHandler) GetAll(c *gin.Context) {
	out, _ := h.svc.FindAll()
	c.JSON(http.StatusOK, out)
}

func (h *CustomerHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	out, err := h.svc.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, out)
}

func (h *CustomerHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var in models.Customer
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	in.ID = uint(id)
	if err := h.svc.Update(&in); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, in)
}
