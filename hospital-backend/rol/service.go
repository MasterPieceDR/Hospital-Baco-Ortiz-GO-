package rol

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAll() ([]Rol, error) {
	return s.repo.GetAll()
}

func (s *Service) GetByID(id uint) (*Rol, error) {
	return s.repo.GetByID(id)
}

func (s *Service) Create(role *Rol) error {
	return s.repo.Create(role)
}

func (s *Service) Update(role *Rol) error {
	return s.repo.Update(role)
}

func (s *Service) Delete(id uint) error {
	return s.repo.Delete(id)
}
