package facturacion

type Service interface {
	CreateFactura(f Factura) error
	GetFactura(id uint) (*Factura, error)
	GetAllFacturas() ([]Factura, error)
	UpdateFactura(f Factura) error
	DeleteFactura(id uint) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) CreateFactura(f Factura) error {
	return s.repo.CreateFactura(f)
}

func (s *service) GetFactura(id uint) (*Factura, error) {
	return s.repo.GetFactura(id)
}

func (s *service) GetAllFacturas() ([]Factura, error) {
	return s.repo.GetAllFacturas()
}

func (s *service) UpdateFactura(f Factura) error {
	return s.repo.UpdateFactura(f)
}

func (s *service) DeleteFactura(id uint) error {
	return s.repo.DeleteFactura(id)
}
