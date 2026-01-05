package prueba

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	service *Service
}

func NewHandler(db *gorm.DB) *Handler {
	repo := NewRepository(db)
	service := NewService(repo)
	return &Handler{service: service}
}

type createUpdateRequest struct {
	PacienteID      uint       `json:"paciente_id" binding:"required"`
	ConsultaID      *uint      `json:"consulta_id"`
	TipoPruebaID    uint       `json:"tipo_prueba_id" binding:"required"`
	FechaSolicitud  time.Time  `json:"fecha_solicitud" binding:"required"`
	FechaResultado  *time.Time `json:"fecha_resultado"`
	Resultado       *string    `json:"resultado"`
	Unidad          *string    `json:"unidad"`
	ValorReferencia *string    `json:"valor_referencia"`
}

func (h *Handler) RegisterRoutes(rg *gin.RouterGroup) {
	pr := rg.Group("/pruebas")

	pr.POST("", h.Create)
	pr.PUT("/:id", h.Update)
	pr.DELETE("/:id", h.Delete)
}

func (h *Handler) GetAll(c *gin.Context) {
	pruebas, err := h.service.Listar()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener las pruebas"})
		return
	}
	c.JSON(http.StatusOK, pruebas)
}

func (h *Handler) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	p, err := h.service.ObtenerPorID(uint(id64))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Prueba no encontrada"})
		return
	}
	c.JSON(http.StatusOK, p)
}

func (h *Handler) GetByPaciente(c *gin.Context) {
	idParam := c.Param("pacienteID")
	id64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de paciente inválido"})
		return
	}
	pruebas, err := h.service.ObtenerPorPaciente(uint(id64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener las pruebas del paciente"})
		return
	}
	c.JSON(http.StatusOK, pruebas)
}

func (h *Handler) Create(c *gin.Context) {
	var req createUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	p := &Prueba{
		PacienteID:      req.PacienteID,
		ConsultaID:      req.ConsultaID,
		TipoPruebaID:    req.TipoPruebaID,
		FechaSolicitud:  req.FechaSolicitud,
		FechaResultado:  req.FechaResultado,
		Resultado:       req.Resultado,
		Unidad:          req.Unidad,
		ValorReferencia: req.ValorReferencia,
	}

	creada, err := h.service.Crear(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear la prueba"})
		return
	}
	c.JSON(http.StatusCreated, creada)
}

func (h *Handler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req createUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	p := &Prueba{
		PacienteID:      req.PacienteID,
		ConsultaID:      req.ConsultaID,
		TipoPruebaID:    req.TipoPruebaID,
		FechaSolicitud:  req.FechaSolicitud,
		FechaResultado:  req.FechaResultado,
		Resultado:       req.Resultado,
		Unidad:          req.Unidad,
		ValorReferencia: req.ValorReferencia,
	}

	actualizada, err := h.service.Actualizar(uint(id64), p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar la prueba"})
		return
	}
	c.JSON(http.StatusOK, actualizada)
}

func (h *Handler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	if err := h.service.Eliminar(uint(id64)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar la prueba"})
		return
	}
	c.Status(http.StatusNoContent)
}
