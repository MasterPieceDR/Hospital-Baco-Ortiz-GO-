package sala

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo}
}

func (s *Service) GetAll() ([]Sala, error) {
	return s.repo.GetAll()
}

func (s *Service) GetByID(id uint) (*Sala, error) {
	return s.repo.GetByID(id)
}

func (s *Service) Create(data *Sala) error {
	return s.repo.Create(data)
}

func (s *Service) Update(data *Sala) error {
	return s.repo.Update(data)
}

func (s *Service) Delete(id uint) error {
	return s.repo.Delete(id)
}
