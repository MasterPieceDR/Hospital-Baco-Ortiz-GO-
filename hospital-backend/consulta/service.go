package consulta

import (
	"errors"
	"time"
)

type ConsultaService interface {
	Crear(data *Consulta) (*Consulta, error)
	Obtener(id uint) (*Consulta, error)
	PorPaciente(id uint) ([]Consulta, error)
	PorMedico(id uint) ([]Consulta, error)
	PorCita(id uint) (*Consulta, error)
	PorRango(desde, hasta time.Time) ([]Consulta, error)
	Actualizar(id uint, data *Consulta) (*Consulta, error)
	Eliminar(id uint) error
}

type consultaService struct {
	repo ConsultaRepository
}

func NewConsultaService(repo ConsultaRepository) ConsultaService {
	return &consultaService{repo: repo}
}

func (s *consultaService) Crear(data *Consulta) (*Consulta, error) {
	if data.Fecha.IsZero() {
		data.Fecha = time.Now()
	}

	err := s.repo.Crear(data)
	return data, err
}

func (s *consultaService) Obtener(id uint) (*Consulta, error) {
	cons, err := s.repo.BuscarPorID(id)
	if err != nil {
		return nil, errors.New("consulta no encontrada")
	}
	return cons, nil
}

func (s *consultaService) PorPaciente(id uint) ([]Consulta, error) {
	return s.repo.BuscarPorPaciente(id)
}

func (s *consultaService) PorMedico(id uint) ([]Consulta, error) {
	return s.repo.BuscarPorMedico(id)
}

func (s *consultaService) PorCita(id uint) (*Consulta, error) {
	return s.repo.BuscarPorCita(id)
}

func (s *consultaService) PorRango(desde, hasta time.Time) ([]Consulta, error) {
	return s.repo.BuscarPorRango(desde, hasta)
}

func (s *consultaService) Actualizar(id uint, data *Consulta) (*Consulta, error) {
	existe, err := s.repo.BuscarPorID(id)
	if err != nil {
		return nil, errors.New("consulta no encontrada")
	}

	existe.MedicoID = data.MedicoID
	existe.PacienteID = data.PacienteID
	if !data.Fecha.IsZero() {
		existe.Fecha = data.Fecha
	}
	existe.Motivo = data.Motivo
	existe.Notas = data.Notas

	err = s.repo.Actualizar(existe)
	return existe, err
}

func (s *consultaService) Eliminar(id uint) error {
	return s.repo.Eliminar(id)
}
