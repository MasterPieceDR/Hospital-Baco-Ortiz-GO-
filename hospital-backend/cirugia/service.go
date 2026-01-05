package cirugia

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Listar() ([]Cirugia, error) {
	return s.repo.GetAll()
}

func (s *Service) ObtenerPorID(id uint) (*Cirugia, error) {
	return s.repo.GetByID(id)
}

func (s *Service) ObtenerPorPaciente(pacienteID uint) ([]Cirugia, error) {
	return s.repo.GetByPaciente(pacienteID)
}

func (s *Service) Crear(c *Cirugia) (*Cirugia, error) {
	if err := s.repo.Create(c); err != nil {
		return nil, err
	}
	return c, nil
}

func (s *Service) Actualizar(id uint, data *Cirugia) (*Cirugia, error) {
	existe, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	existe.HospitalizacionID = data.HospitalizacionID
	existe.PacienteID = data.PacienteID
	existe.MedicoID = data.MedicoID
	existe.AnestesistaID = data.AnestesistaID
	existe.FechaCirugia = data.FechaCirugia
	existe.SalaID = data.SalaID
	existe.Tipo = data.Tipo
	existe.Notas = data.Notas

	if err := s.repo.Update(existe); err != nil {
		return nil, err
	}

	return existe, nil
}

func (s *Service) Eliminar(id uint) error {
	return s.repo.Delete(id)
}
