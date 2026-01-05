package receta

type Service interface {
	GetAll() ([]Receta, error)
	GetByID(id uint) (*Receta, error)
	Create(r *Receta) error
	Update(r *Receta) error
	Delete(id uint) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetAll() ([]Receta, error) {
	return s.repo.GetAll()
}

func (s *service) GetByID(id uint) (*Receta, error) {
	return s.repo.GetByID(id)
}

func (s *service) Create(rec *Receta) error {
	return s.repo.Create(rec)
}

func (s *service) Update(rec *Receta) error {
	return s.repo.Update(rec)
}

func (s *service) Delete(id uint) error {
	return s.repo.Delete(id)
}
