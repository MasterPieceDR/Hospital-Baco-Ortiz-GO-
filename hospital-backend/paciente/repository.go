package paciente

import "gorm.io/gorm"

type PacienteRepository interface {
	Crear(paciente *Paciente) error
	BuscarPorID(id uint) (*Paciente, error)
	BuscarTodos() ([]Paciente, error)
	Actualizar(paciente *Paciente) error
	Eliminar(id uint) error
}

type pacienteRepository struct {
	db *gorm.DB
}

func NewPacienteRepository(db *gorm.DB) PacienteRepository {
	return &pacienteRepository{db}
}

func (r *pacienteRepository) Crear(paciente *Paciente) error {
	return r.db.Create(paciente).Error
}

func (r *pacienteRepository) BuscarPorID(id uint) (*Paciente, error) {
	var paciente Paciente
	err := r.db.First(&paciente, id).Error
	return &paciente, err
}

func (r *pacienteRepository) BuscarTodos() ([]Paciente, error) {
	var pacientes []Paciente
	err := r.db.Find(&pacientes).Error
	return pacientes, err
}

func (r *pacienteRepository) Actualizar(paciente *Paciente) error {
	return r.db.Save(paciente).Error
}

func (r *pacienteRepository) Eliminar(id uint) error {
	return r.db.Delete(&Paciente{}, id).Error
}
