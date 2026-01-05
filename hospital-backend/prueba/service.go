package prueba

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Listar() ([]Prueba, error) {
	return s.repo.GetAll()
}

func (s *Service) ObtenerPorID(id uint) (*Prueba, error) {
	return s.repo.GetByID(id)
}

func (s *Service) ObtenerPorPaciente(pacienteID uint) ([]Prueba, error) {
	return s.repo.GetByPacienteID(pacienteID)
}

func (s *Service) Crear(p *Prueba) (*Prueba, error) {
	if err := s.repo.Create(p); err != nil {
		return nil, err
	}
	return p, nil
}

func (s *Service) Actualizar(id uint, datos *Prueba) (*Prueba, error) {
	existente, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	existente.PacienteID = datos.PacienteID
	existente.ConsultaID = datos.ConsultaID
	existente.TipoPruebaID = datos.TipoPruebaID
	existente.FechaSolicitud = datos.FechaSolicitud
	existente.FechaResultado = datos.FechaResultado
	existente.Resultado = datos.Resultado
	existente.Unidad = datos.Unidad
	existente.ValorReferencia = datos.ValorReferencia

	if err := s.repo.Update(existente); err != nil {
		return nil, err
	}
	return existente, nil
}

func (s *Service) Eliminar(id uint) error {
	return s.repo.Delete(id)
}
