package inventario

type Service interface {
	GetAll() ([]Inventario, error)
	GetByID(id uint) (*Inventario, error)
	Create(inv *Inventario) error
	Update(inv *Inventario) error
	Delete(id uint) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetAll() ([]Inventario, error) {
	return s.repo.GetAll()
}

func (s *service) GetByID(id uint) (*Inventario, error) {
	return s.repo.GetByID(id)
}

func (s *service) Create(inv *Inventario) error {
	return s.repo.Create(inv)
}

func (s *service) Update(inv *Inventario) error {
	return s.repo.Update(inv)
}

func (s *service) Delete(id uint) error {
	return s.repo.Delete(id)
}
