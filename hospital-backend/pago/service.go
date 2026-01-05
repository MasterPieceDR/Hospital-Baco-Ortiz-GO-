package pago

import "errors"

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) GetAll() ([]Pago, error) {
	return s.repo.FindAll()
}

func (s *Service) GetByID(id uint) (*Pago, error) {
	return s.repo.FindByID(id)
}

func (s *Service) Create(p *Pago) error {
	if p.Monto <= 0 {
		return errors.New("el monto debe ser mayor a cero")
	}
	return s.repo.Create(p)
}

func (s *Service) Update(p *Pago) error {
	return s.repo.Update(p)
}

func (s *Service) Delete(id uint) error {
	return s.repo.Delete(id)
}
