package medico

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo}
}

func (s *Service) GetAll() ([]Medico, error) {
	return s.repo.FindAll()
}

func (s *Service) GetByID(id uint) (*Medico, error) {
	return s.repo.FindByID(id)
}

func (s *Service) Create(m *Medico) error {
	return s.repo.Create(m)
}

func (s *Service) Update(m *Medico) error {
	return s.repo.Update(m)
}

func (s *Service) Delete(id uint) error {
	return s.repo.Delete(id)
}
