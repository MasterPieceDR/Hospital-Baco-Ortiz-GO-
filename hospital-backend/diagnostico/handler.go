package diagnostico

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
	return &Handler{
		service: NewService(NewRepository(db)),
	}
}

type createUpdateRequest struct {
	ConsultaID  uint      `json:"consulta_id" binding:"required"`
	Descripcion string    `json:"descripcion" binding:"required"`
	Fecha       time.Time `json:"fecha" binding:"required"`
}

func (h *Handler) RegisterRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/diagnosticos")

	r.GET("", h.GetAll)
	r.GET("/:id", h.GetByID)
	r.GET("/consulta/:consultaID", h.GetByConsulta)
	r.POST("", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
}

func (h *Handler) GetAll(c *gin.Context) {
	lista, err := h.service.Listar()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo diagnósticos"})
		return
	}
	c.JSON(http.StatusOK, lista)
}

func (h *Handler) GetByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	item, err := h.service.ObtenerPorID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Diagnóstico no encontrado"})
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetByConsulta(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("consultaID"), 10, 32)
	lista, err := h.service.ObtenerPorConsulta(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo diagnósticos"})
		return
	}
	c.JSON(http.StatusOK, lista)
}

func (h *Handler) Create(c *gin.Context) {
	var req createUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	n := &Diagnostico{
		ConsultaID:  req.ConsultaID,
		Descripcion: req.Descripcion,
		Fecha:       req.Fecha,
	}

	creado, err := h.service.Crear(n)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear"})
		return
	}
	c.JSON(http.StatusCreated, creado)
}

func (h *Handler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	var req createUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	n := &Diagnostico{
		ConsultaID:  req.ConsultaID,
		Descripcion: req.Descripcion,
		Fecha:       req.Fecha,
	}

	actualizado, err := h.service.Actualizar(uint(id), n)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error actualizando"})
		return
	}

	c.JSON(http.StatusOK, actualizado)
}

func (h *Handler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if err := h.service.Eliminar(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error eliminando"})
		return
	}

	c.Status(http.StatusNoContent)
}
