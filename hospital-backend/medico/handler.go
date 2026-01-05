package medico

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service}
}

func (h *Handler) GetAll(c *gin.Context) {
	medicos, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener los médicos"})
		return
	}
	c.JSON(http.StatusOK, medicos)
}

func (h *Handler) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	medico, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Médico no encontrado"})
		return
	}
	c.JSON(http.StatusOK, medico)
}

func (h *Handler) Create(c *gin.Context) {
	var m Medico
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	if err := h.service.Create(&m); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear el médico"})
		return
	}

	c.JSON(http.StatusCreated, m)
}

func (h *Handler) Update(c *gin.Context) {
	var m Medico
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	if err := h.service.Update(&m); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar"})
		return
	}

	c.JSON(http.StatusOK, m)
}

func (h *Handler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Eliminado correctamente"})
}
