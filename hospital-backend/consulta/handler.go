package consulta

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ConsultaHandler struct {
	service ConsultaService
}

func NewConsultaHandler(service ConsultaService) *ConsultaHandler {
	return &ConsultaHandler{service: service}
}

func (h *ConsultaHandler) Crear(c *gin.Context) {
	var data Consulta
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	res, err := h.service.Crear(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h *ConsultaHandler) Obtener(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	res, err := h.service.Obtener(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Consulta no encontrada"})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *ConsultaHandler) PorPaciente(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("paciente_id"))
	res, err := h.service.PorPaciente(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *ConsultaHandler) PorMedico(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("medico_id"))
	res, err := h.service.PorMedico(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *ConsultaHandler) PorCita(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("cita_id"))
	res, err := h.service.PorCita(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No existe consulta para la cita"})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *ConsultaHandler) PorRango(c *gin.Context) {
	layout := "2006-01-02"
	desdeStr := c.Query("desde")
	hastaStr := c.Query("hasta")

	desde, err1 := time.Parse(layout, desdeStr)
	hasta, err2 := time.Parse(layout, hastaStr)

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato incorrecto. Formato correcto: 2006-01-02"})
		return
	}

	res, err := h.service.PorRango(desde, hasta)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *ConsultaHandler) Actualizar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var data Consulta
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	res, err := h.service.Actualizar(uint(id), &data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *ConsultaHandler) Eliminar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.service.Eliminar(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Consulta eliminada"})
}
