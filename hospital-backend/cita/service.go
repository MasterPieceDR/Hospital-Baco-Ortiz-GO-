package cita

import (
	"errors"
	"time"
)

type CitaService interface {
	CrearCita(data *Cita) (*Cita, error)
	ObtenerCita(id uint) (*Cita, error)
	ListarCitas() ([]Cita, error)
	ListarCitasPorMedico(id uint) ([]Cita, error)
	ListarCitasPorPaciente(id uint) ([]Cita, error)
	ListarCitasPorRango(desde, hasta time.Time) ([]Cita, error)
	ActualizarCita(id uint, data *Cita) (*Cita, error)
	CancelarCita(id uint) (*Cita, error)
	EliminarCita(id uint) error
}

type citaService struct {
	repo CitaRepository
}

func NewCitaService(repo CitaRepository) CitaService {
	return &citaService{repo: repo}
}

func (s *citaService) CrearCita(data *Cita) (*Cita, error) {
	if data.Estado == "" {
		data.Estado = "Pendiente"
	}
	err := s.repo.Crear(data)
	return data, err
}

func (s *citaService) ObtenerCita(id uint) (*Cita, error) {
	c, err := s.repo.BuscarPorID(id)
	if err != nil {
		return nil, errors.New("cita no encontrada")
	}
	return c, nil
}

func (s *citaService) ListarCitas() ([]Cita, error) {
	return s.repo.BuscarTodas()
}

func (s *citaService) ListarCitasPorMedico(id uint) ([]Cita, error) {
	return s.repo.BuscarPorMedico(id)
}

func (s *citaService) ListarCitasPorPaciente(id uint) ([]Cita, error) {
	return s.repo.BuscarPorPaciente(id)
}

func (s *citaService) ListarCitasPorRango(desde, hasta time.Time) ([]Cita, error) {
	return s.repo.BuscarPorRangoFecha(desde, hasta)
}

func (s *citaService) ActualizarCita(id uint, data *Cita) (*Cita, error) {
	existe, err := s.repo.BuscarPorID(id)
	if err != nil {
		return nil, errors.New("cita no encontrada")
	}

	existe.PacienteID = data.PacienteID
	existe.MedicoID = data.MedicoID
	existe.SalaID = data.SalaID
	existe.FechaHora = data.FechaHora
	existe.Estado = data.Estado
	existe.Motivo = data.Motivo
	existe.Notas = data.Notas

	err = s.repo.Actualizar(existe)
	return existe, err
}

func (s *citaService) CancelarCita(id uint) (*Cita, error) {
	existe, err := s.repo.BuscarPorID(id)
	if err != nil {
		return nil, errors.New("cita no encontrada")
	}

	existe.Estado = "Cancelada"
	err = s.repo.Actualizar(existe)
	return existe, err
}

func (s *citaService) EliminarCita(id uint) error {
	return s.repo.Eliminar(id)
}
