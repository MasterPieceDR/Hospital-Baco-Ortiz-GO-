package aseguradora

type Service interface {
	GetAll() ([]Aseguradora, error)
	GetByID(id uint) (*Aseguradora, error)
	Create(a *Aseguradora) error
	Update(a *Aseguradora) error
	Delete(id uint) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetAll() ([]Aseguradora, error) {
	return s.repo.GetAll()
}

func (s *service) GetByID(id uint) (*Aseguradora, error) {
	return s.repo.GetByID(id)
}

func (s *service) Create(a *Aseguradora) error {
	return s.repo.Create(a)
}

func (s *service) Update(a *Aseguradora) error {
	return s.repo.Update(a)
}

func (s *service) Delete(id uint) error {
	return s.repo.Delete(id)
}
