package medicamento

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo}
}

func (s *Service) GetAll() ([]Medicamento, error) {
	return s.repo.GetAll()
}

func (s *Service) GetByID(id uint) (*Medicamento, error) {
	return s.repo.GetByID(id)
}

func (s *Service) Create(m *Medicamento) error {
	return s.repo.Create(m)
}

func (s *Service) Update(m *Medicamento) error {
	return s.repo.Update(m)
}

func (s *Service) Delete(id uint) error {
	return s.repo.Delete(id)
}
