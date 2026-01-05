package cita

import (
	"time"

	"gorm.io/gorm"
)

type CitaRepository interface {
	Crear(c *Cita) error
	BuscarPorID(id uint) (*Cita, error)
	BuscarTodas() ([]Cita, error)
	BuscarPorMedico(medicoID uint) ([]Cita, error)
	BuscarPorPaciente(pacienteID uint) ([]Cita, error)
	BuscarPorRangoFecha(desde, hasta time.Time) ([]Cita, error)
	Actualizar(c *Cita) error
	Eliminar(id uint) error
}

type citaRepository struct {
	db *gorm.DB
}

func NewCitaRepository(db *gorm.DB) CitaRepository {
	return &citaRepository{db: db}
}

func (r *citaRepository) Crear(c *Cita) error {
	return r.db.Create(c).Error
}

func (r *citaRepository) BuscarPorID(id uint) (*Cita, error) {
	var cita Cita
	err := r.db.First(&cita, "cita_id = ?", id).Error
	return &cita, err
}

func (r *citaRepository) BuscarTodas() ([]Cita, error) {
	var citas []Cita
	err := r.db.Order("cita_id").Find(&citas).Error
	return citas, err
}

func (r *citaRepository) BuscarPorMedico(id uint) ([]Cita, error) {
	var citas []Cita
	err := r.db.Where("medico_id = ?", id).Find(&citas).Error
	return citas, err
}

func (r *citaRepository) BuscarPorPaciente(id uint) ([]Cita, error) {
	var citas []Cita
	err := r.db.Where("paciente_id = ?", id).Find(&citas).Error
	return citas, err
}

func (r *citaRepository) BuscarPorRangoFecha(desde, hasta time.Time) ([]Cita, error) {
	var citas []Cita
	err := r.db.Where("fecha_hora BETWEEN ? AND ?", desde, hasta).Find(&citas).Error
	return citas, err
}

func (r *citaRepository) Actualizar(c *Cita) error {
	return r.db.Save(c).Error
}

func (r *citaRepository) Eliminar(id uint) error {
	return r.db.Delete(&Cita{}, "cita_id = ?", id).Error
}
