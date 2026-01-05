package tipo_prueba

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Listar() ([]TipoPrueba, error) {
	return s.repo.GetAll()
}

func (s *Service) ObtenerPorID(id int) (*TipoPrueba, error) {
	return s.repo.GetByID(id)
}

func (s *Service) Crear(nombre string, descripcion *string) (*TipoPrueba, error) {
	t := &TipoPrueba{
		Nombre:      nombre,
		Descripcion: descripcion,
	}
	if err := s.repo.Create(t); err != nil {
		return nil, err
	}
	return t, nil
}

func (s *Service) Actualizar(id int, nombre string, descripcion *string) (*TipoPrueba, error) {
	t, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	t.Nombre = nombre
	t.Descripcion = descripcion
	if err := s.repo.Update(t); err != nil {
		return nil, err
	}
	return t, nil
}

func (s *Service) Eliminar(id int) error {
	return s.repo.Delete(id)
}
