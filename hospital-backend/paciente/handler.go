package paciente

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PacienteHandler struct {
	service PacienteService
}

func NewPacienteHandler(service PacienteService) *PacienteHandler {
	return &PacienteHandler{service}
}

func (h *PacienteHandler) Crear(c *gin.Context) {
	var data Paciente
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	paciente, err := h.service.CrearPaciente(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, paciente)
}

func (h *PacienteHandler) ObtenerTodos(c *gin.Context) {
	pacientes, err := h.service.ObtenerPacientes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pacientes)
}

func (h *PacienteHandler) ObtenerPorID(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	paciente, err := h.service.ObtenerPaciente(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Paciente no encontrado"})
		return
	}

	c.JSON(http.StatusOK, paciente)
}

func (h *PacienteHandler) Actualizar(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var data Paciente
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	paciente, err := h.service.ActualizarPaciente(uint(id), &data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, paciente)
}

func (h *PacienteHandler) Eliminar(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	err := h.service.EliminarPaciente(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar paciente"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Paciente eliminado"})
}
