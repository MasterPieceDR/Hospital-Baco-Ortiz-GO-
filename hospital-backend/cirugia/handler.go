package cirugia

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
	HospitalizacionID uint      `json:"hospitalizacion_id"`
	PacienteID        uint      `json:"paciente_id" binding:"required"`
	MedicoID          uint      `json:"medico_id" binding:"required"`
	AnestesistaID     *uint     `json:"anestesista_id"`
	FechaCirugia      time.Time `json:"fecha_cirugia" binding:"required"`
	SalaID            uint      `json:"sala_id"`
	Tipo              string    `json:"tipo"`
	Notas             string    `json:"notas"`
}

func (h *Handler) RegisterRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/cirugias")

	r.GET("", h.GetAll)
	r.GET("/:id", h.GetByID)
	r.GET("/paciente/:id", h.GetByPaciente)
	r.POST("", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
}

func (h *Handler) GetAll(c *gin.Context) {
	lista, err := h.service.Listar()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo cirugías"})
		return
	}
	c.JSON(http.StatusOK, lista)
}

func (h *Handler) GetByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	cirugia, err := h.service.ObtenerPorID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cirugía no encontrada"})
		return
	}
	c.JSON(http.StatusOK, cirugia)
}

func (h *Handler) GetByPaciente(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	lista, err := h.service.ObtenerPorPaciente(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo cirugías"})
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

	cirugia := &Cirugia{
		HospitalizacionID: req.HospitalizacionID,
		PacienteID:        req.PacienteID,
		MedicoID:          req.MedicoID,
		AnestesistaID:     req.AnestesistaID,
		FechaCirugia:      req.FechaCirugia,
		SalaID:            req.SalaID,
		Tipo:              req.Tipo,
		Notas:             req.Notas,
	}

	creada, err := h.service.Crear(cirugia)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creando cirugía"})
		return
	}

	c.JSON(http.StatusCreated, creada)
}

func (h *Handler) Update(c *gin.Context) {
	id64, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	id := uint(id64)

	var req createUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	data := &Cirugia{
		CirugiaID:         id,
		HospitalizacionID: req.HospitalizacionID,
		PacienteID:        req.PacienteID,
		MedicoID:          req.MedicoID,
		AnestesistaID:     req.AnestesistaID,
		FechaCirugia:      req.FechaCirugia,
		SalaID:            req.SalaID,
		Tipo:              req.Tipo,
		Notas:             req.Notas,
	}

	actualizada, err := h.service.Actualizar(id, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error actualizando cirugía"})
		return
	}

	c.JSON(http.StatusOK, actualizada)
}

func (h *Handler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if err := h.service.Eliminar(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error eliminando cirugía"})
		return
	}

	c.Status(http.StatusNoContent)
}
