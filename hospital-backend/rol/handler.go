package rol

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

// Esta función la usaremos luego desde routes.RegisterRoutes(...)
func (h *Handler) RegisterRoutes(rg *gin.RouterGroup) {
	roles := rg.Group("/roles")

	roles.GET("/", h.GetAll)
	roles.GET("/:id", h.GetByID)
	roles.POST("/", h.Create)
	roles.PUT("/:id", h.Update)
	roles.DELETE("/:id", h.Delete)
}

func (h *Handler) GetAll(c *gin.Context) {
	roles, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener los roles"})
		return
	}
	c.JSON(http.StatusOK, roles)
}

func parseUintParam(c *gin.Context, name string) (uint, bool) {
	idStr := c.Param(name)
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return 0, false
	}
	return uint(id64), true
}

func (h *Handler) GetByID(c *gin.Context) {
	id, ok := parseUintParam(c, "id")
	if !ok {
		return
	}

	role, err := h.service.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Rol no encontrado"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el rol"})
		}
		return
	}

	c.JSON(http.StatusOK, role)
}

func (h *Handler) Create(c *gin.Context) {
	var input Rol
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	input.RolID = 0

	if err := h.service.Create(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el rol"})
		return
	}

	c.JSON(http.StatusCreated, input)
}

func (h *Handler) Update(c *gin.Context) {
	id, ok := parseUintParam(c, "id")
	if !ok {
		return
	}

	var input Rol
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	role, err := h.service.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Rol no encontrado"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar el rol"})
		}
		return
	}

	role.Nombre = input.Nombre
	role.Descripcion = input.Descripcion

	if err := h.service.Update(role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el rol"})
		return
	}

	c.JSON(http.StatusOK, role)
}

func (h *Handler) Delete(c *gin.Context) {
	id, ok := parseUintParam(c, "id")
	if !ok {
		return
	}

	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el rol"})
		return
	}

	c.Status(http.StatusNoContent)
}
