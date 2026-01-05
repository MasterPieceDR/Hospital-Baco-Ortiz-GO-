package paciente

import "errors"

type PacienteService interface {
	CrearPaciente(data *Paciente) (*Paciente, error)
	ObtenerPaciente(id uint) (*Paciente, error)
	ObtenerPacientes() ([]Paciente, error)
	ActualizarPaciente(id uint, data *Paciente) (*Paciente, error)
	EliminarPaciente(id uint) error
}

type pacienteService struct {
	repo PacienteRepository
}

func NewPacienteService(repo PacienteRepository) PacienteService {
	return &pacienteService{repo}
}

func (s *pacienteService) CrearPaciente(data *Paciente) (*Paciente, error) {

	data.Estado = true

	err := s.repo.Crear(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *pacienteService) ObtenerPaciente(id uint) (*Paciente, error) {
	paciente, err := s.repo.BuscarPorID(id)
	if err != nil {
		return nil, errors.New("paciente no encontrado")
	}
	return paciente, nil
}

func (s *pacienteService) ObtenerPacientes() ([]Paciente, error) {
	return s.repo.BuscarTodos()
}

func (s *pacienteService) ActualizarPaciente(id uint, data *Paciente) (*Paciente, error) {

	existe, err := s.repo.BuscarPorID(id)
	if err != nil {
		return nil, errors.New("paciente no encontrado")
	}

	existe.Nombre = data.Nombre
	existe.Apellido = data.Apellido
	existe.Cedula = data.Cedula
	existe.Correo = data.Correo
	existe.Telefono = data.Telefono
	existe.Direccion = data.Direccion
	existe.FechaNacimiento = data.FechaNacimiento
	existe.Estado = data.Estado

	err = s.repo.Actualizar(existe)
	if err != nil {
		return nil, err
	}

	return existe, nil
}

func (s *pacienteService) EliminarPaciente(id uint) error {
	return s.repo.Eliminar(id)
}
