package usuario

import (
	"hospital-backend/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token      string   `json:"token"`
	Username   string   `json:"username"`
	Roles      []string `json:"roles"`
	MedicoID   *int     `json:"medico_id,omitempty"`
	PacienteID *int     `json:"paciente_id,omitempty"`
}

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	u, roles, err := h.service.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := auth.GenerateJWT(u.UsuarioID, u.Username, roles, u.MedicoIDAsociado, u.PacienteIDAsociado)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error generando token"})
		return
	}

	resp := LoginResponse{
		Token:      token,
		Username:   u.Username,
		Roles:      roles,
		MedicoID:   u.MedicoIDAsociado,
		PacienteID: u.PacienteIDAsociado,
	}

	c.JSON(http.StatusOK, resp)
}
