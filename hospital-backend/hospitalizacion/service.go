package hospitalizacion

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo}
}

func (s *Service) GetAll() ([]Hospitalizacion, error) {
	return s.repo.GetAll()
}

func (s *Service) GetByID(id uint) (*Hospitalizacion, error) {
	return s.repo.GetByID(id)
}

func (s *Service) Create(h *Hospitalizacion) error {
	return s.repo.Create(h)
}

func (s *Service) Update(h *Hospitalizacion) error {
	return s.repo.Update(h)
}

func (s *Service) Delete(id uint) error {
	return s.repo.Delete(id)
}
