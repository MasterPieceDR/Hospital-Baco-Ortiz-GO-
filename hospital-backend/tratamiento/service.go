package tratamiento

type Service interface {
	GetAll() ([]Tratamiento, error)
	GetByID(id int) (*Tratamiento, error)
	Create(t *Tratamiento) error
	Update(t *Tratamiento) error
	Delete(id int) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetAll() ([]Tratamiento, error) {
	return s.repo.GetAll()
}

func (s *service) GetByID(id int) (*Tratamiento, error) {
	return s.repo.GetByID(id)
}

func (s *service) Create(t *Tratamiento) error {
	return s.repo.Create(t)
}

func (s *service) Update(t *Tratamiento) error {
	return s.repo.Update(t)
}

func (s *service) Delete(id int) error {
	return s.repo.Delete(id)
}
