package facturacion

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{s}
}

func (h *Handler) CreateFactura(c *gin.Context) {
	var f Factura
	if err := c.ShouldBindJSON(&f); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateFactura(f); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, f)
}

func (h *Handler) GetFactura(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	factura, err := h.service.GetFactura(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Factura no encontrada"})
		return
	}

	c.JSON(http.StatusOK, factura)
}

func (h *Handler) GetAllFacturas(c *gin.Context) {
	facturas, err := h.service.GetAllFacturas()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, facturas)
}

func (h *Handler) UpdateFactura(c *gin.Context) {
	var f Factura

	if err := c.ShouldBindJSON(&f); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateFactura(f); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, f)
}

func (h *Handler) DeleteFactura(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.service.DeleteFactura(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Factura eliminada"})
}
