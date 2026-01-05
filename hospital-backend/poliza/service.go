package poliza

type Service interface {
	GetAll() ([]Poliza, error)
	GetByID(id uint) (*Poliza, error)
	Create(p *Poliza) error
	Update(p *Poliza) error
	Delete(id uint) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetAll() ([]Poliza, error) {
	return s.repo.GetAll()
}

func (s *service) GetByID(id uint) (*Poliza, error) {
	return s.repo.GetByID(id)
}

func (s *service) Create(p *Poliza) error {
	return s.repo.Create(p)
}

func (s *service) Update(p *Poliza) error {
	return s.repo.Update(p)
}

func (s *service) Delete(id uint) error {
	return s.repo.Delete(id)
}
