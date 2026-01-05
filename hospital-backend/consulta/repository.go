package consulta

import (
	"time"

	"gorm.io/gorm"
)

type ConsultaRepository interface {
	Crear(c *Consulta) error
	BuscarPorID(id uint) (*Consulta, error)
	BuscarPorPaciente(id uint) ([]Consulta, error)
	BuscarPorMedico(id uint) ([]Consulta, error)
	BuscarPorCita(id uint) (*Consulta, error)
	BuscarPorRango(desde, hasta time.Time) ([]Consulta, error)
	Actualizar(c *Consulta) error
	Eliminar(id uint) error
}

type consultaRepository struct {
	db *gorm.DB
}

func NewConsultaRepository(db *gorm.DB) ConsultaRepository {
	return &consultaRepository{db: db}
}

func (r *consultaRepository) Crear(c *Consulta) error {
	return r.db.Create(c).Error
}

func (r *consultaRepository) BuscarPorID(id uint) (*Consulta, error) {
	var cons Consulta
	err := r.db.Preload("Diagnosticos").First(&cons, "consulta_id = ?", id).Error
	return &cons, err
}

func (r *consultaRepository) BuscarPorPaciente(id uint) ([]Consulta, error) {
	var cons []Consulta
	err := r.db.Preload("Diagnosticos").
		Where("paciente_id = ?", id).
		Order("fecha_consulta DESC").
		Find(&cons).Error
	return cons, err
}

func (r *consultaRepository) BuscarPorMedico(id uint) ([]Consulta, error) {
	var cons []Consulta
	err := r.db.Preload("Diagnosticos").
		Where("medico_id = ?", id).
		Order("fecha_consulta DESC").
		Find(&cons).Error
	return cons, err
}

func (r *consultaRepository) BuscarPorCita(id uint) (*Consulta, error) {
	var cons Consulta
	err := r.db.Preload("Diagnosticos").Where("cita_id = ?", id).First(&cons).Error
	return &cons, err
}

func (r *consultaRepository) BuscarPorRango(desde, hasta time.Time) ([]Consulta, error) {
	var cons []Consulta
	err := r.db.Preload("Diagnosticos").
		Where("fecha_consulta BETWEEN ? AND ?", desde, hasta).
		Order("fecha_consulta DESC").
		Find(&cons).Error
	return cons, err
}

func (r *consultaRepository) Actualizar(c *Consulta) error {
	return r.db.Save(c).Error
}

func (r *consultaRepository) Eliminar(id uint) error {
	return r.db.Delete(&Consulta{}, "consulta_id = ?", id).Error
}
