package usuario_rol

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetRolesByUsuarioID(usuarioID int) ([]int, error) {
	return s.repo.GetRolesByUsuarioID(usuarioID)
}

func (s *Service) AddRole(usuarioID, rolID int) error {
	return s.repo.AddRole(usuarioID, rolID)
}

func (s *Service) RemoveRole(usuarioID, rolID int) error {
	return s.repo.RemoveRole(usuarioID, rolID)
}
