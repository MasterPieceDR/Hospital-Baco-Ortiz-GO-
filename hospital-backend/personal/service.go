package personal

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo}
}

func (s *Service) GetAll() ([]Personal, error) {
	return s.repo.GetAll()
}

func (s *Service) GetByID(id uint) (*Personal, error) {
	return s.repo.GetByID(id)
}

func (s *Service) Create(p *Personal) error {
	return s.repo.Create(p)
}

func (s *Service) Update(p *Personal) error {
	return s.repo.Update(p)
}

func (s *Service) Delete(id uint) error {
	return s.repo.Delete(id)
}
