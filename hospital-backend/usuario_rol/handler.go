package usuario_rol

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

func (h *Handler) RegisterRoutes(rg *gin.RouterGroup) {
	usrRol := rg.Group("/usuario-rol")
	usrRol.GET("/:usuarioID", h.GetRoles)
	usrRol.POST("", h.AddRole)
	usrRol.DELETE("/:usuarioID/:rolID", h.DeleteRole)
}

func (h *Handler) GetRoles(c *gin.Context) {
	usuarioID, err := strconv.Atoi(c.Param("usuarioID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}
	roles, err := h.service.GetRolesByUsuarioID(usuarioID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error obteniendo roles"})
		return
	}
	c.JSON(http.StatusOK, roles)
}

func (h *Handler) AddRole(c *gin.Context) {
	var req struct {
		UsuarioID int `json:"usuario_id"`
		RolID     int `json:"rol_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "json inválido"})
		return
	}
	if err := h.service.AddRole(req.UsuarioID, req.RolID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error asignando rol"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "rol asignado"})
}

func (h *Handler) DeleteRole(c *gin.Context) {
	usuarioID, err := strconv.Atoi(c.Param("usuarioID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}
	rolID, err := strconv.Atoi(c.Param("rolID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}
	if err := h.service.RemoveRole(usuarioID, rolID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error eliminando rol"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "rol eliminado"})
}
