package diagnostico

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Listar() ([]Diagnostico, error) {
	return s.repo.GetAll()
}

func (s *Service) ObtenerPorID(id uint) (*Diagnostico, error) {
	return s.repo.GetByID(id)
}

func (s *Service) ObtenerPorConsulta(consultaID uint) ([]Diagnostico, error) {
	return s.repo.GetByConsulta(consultaID)
}

func (s *Service) Crear(d *Diagnostico) (*Diagnostico, error) {
	if err := s.repo.Create(d); err != nil {
		return nil, err
	}
	return d, nil
}

func (s *Service) Actualizar(id uint, data *Diagnostico) (*Diagnostico, error) {
	exist, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	exist.ConsultaID = data.ConsultaID
	exist.Descripcion = data.Descripcion
	exist.Fecha = data.Fecha

	if err := s.repo.Update(exist); err != nil {
		return nil, err
	}
	return exist, nil
}

func (s *Service) Eliminar(id uint) error {
	return s.repo.Delete(id)
}
