package tipo_prueba

import (
	"net/http"
	"strconv"

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
	Nombre      string  `json:"nombre" binding:"required"`
	Descripcion *string `json:"descripcion"`
}

func (h *Handler) RegisterRoutes(rg *gin.RouterGroup) {
	tipoPrueba := rg.Group("/tipos-prueba")
	tipoPrueba.GET("", h.GetAll)
	tipoPrueba.GET("/:id", h.GetByID)
	tipoPrueba.POST("", h.Create)
	tipoPrueba.PUT("/:id", h.Update)
	tipoPrueba.DELETE("/:id", h.Delete)
}

func (h *Handler) GetAll(c *gin.Context) {
	tipos, err := h.service.Listar()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudieron obtener los tipos de prueba"})
		return
	}
	c.JSON(http.StatusOK, tipos)
}

func (h *Handler) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}
	tipo, err := h.service.ObtenerPorID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "tipo de prueba no encontrado"})
		return
	}
	c.JSON(http.StatusOK, tipo)
}

func (h *Handler) Create(c *gin.Context) {
	var req createUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "json inválido"})
		return
	}
	tipo, err := h.service.Crear(req.Nombre, req.Descripcion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudo crear el tipo de prueba"})
		return
	}
	c.JSON(http.StatusCreated, tipo)
}

func (h *Handler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}
	var req createUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "json inválido"})
		return
	}
	tipo, err := h.service.Actualizar(id, req.Nombre, req.Descripcion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudo actualizar el tipo de prueba"})
		return
	}
	c.JSON(http.StatusOK, tipo)
}

func (h *Handler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}
	if err := h.service.Eliminar(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudo eliminar el tipo de prueba"})
		return
	}
	c.Status(http.StatusNoContent)
}
