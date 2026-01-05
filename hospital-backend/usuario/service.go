package usuario

import (
	"errors"
	"hospital-backend/auth"

	"gorm.io/gorm"
)

type Service struct {
	repo *Repository
}

func NewService(db *gorm.DB) *Service {
	return &Service{repo: NewRepository(db)}
}

func (s *Service) Login(username, password string) (*Usuario, []string, error) {
	u, err := s.repo.FindByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, errors.New("usuario o contraseña incorrectos")
		}
		return nil, nil, err
	}

	if !u.Activo {
		return nil, nil, errors.New("usuario inactivo")
	}

	if !auth.CheckPassword(u.PasswordHash, password) {
		return nil, nil, errors.New("usuario o contraseña incorrectos")
	}

	roles, err := s.repo.GetRolesByUsuarioID(u.UsuarioID)
	if err != nil {
		return nil, nil, err
	}

	return u, roles, nil
}
