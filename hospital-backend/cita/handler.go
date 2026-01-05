package cita

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CitaHandler struct {
	service CitaService
}

func NewCitaHandler(service CitaService) *CitaHandler {
	return &CitaHandler{service: service}
}

func (h *CitaHandler) Registrar(rg *gin.RouterGroup) {
	r := rg.Group("/citas")

	r.POST("", h.Crear)
	r.GET("", h.ObtenerTodas)
	r.GET("/:id", h.ObtenerPorID)
	r.GET("/medico/:id", h.ObtenerPorMedico)
	r.GET("/paciente/:id", h.ObtenerPorPaciente)
	r.GET("/rango", h.ObtenerPorRangoFechas)
	r.PUT("/:id", h.Actualizar)
	r.PUT("/:id/cancelar", h.Cancelar)
	r.DELETE("/:id", h.Eliminar)
}

func (h *CitaHandler) Crear(c *gin.Context) {
	var data Cita
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	creada, err := h.service.CrearCita(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, creada)
}

func (h *CitaHandler) ObtenerTodas(c *gin.Context) {
	list, err := h.service.ListarCitas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *CitaHandler) ObtenerPorID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	item, err := h.service.ObtenerCita(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cita no encontrada"})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *CitaHandler) ObtenerPorMedico(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	list, err := h.service.ListarCitasPorMedico(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *CitaHandler) ObtenerPorPaciente(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	list, err := h.service.ListarCitasPorPaciente(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *CitaHandler) ObtenerPorRangoFechas(c *gin.Context) {
	layout := "2006-01-02 15:04"
	desdeStr := c.Query("desde")
	hastaStr := c.Query("hasta")

	desde, err1 := time.Parse(layout, desdeStr)
	hasta, err2 := time.Parse(layout, hastaStr)

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Usa formato: 2006-01-02 15:04"})
		return
	}

	list, err := h.service.ListarCitasPorRango(desde, hasta)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *CitaHandler) Actualizar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var data Cita
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.service.ActualizarCita(uint(id), &data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *CitaHandler) Cancelar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	item, err := h.service.CancelarCita(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *CitaHandler) Eliminar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.service.EliminarCita(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error eliminando cita"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Cita eliminada"})
}
