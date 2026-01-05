package hospitalizacion

import "time"

type Hospitalizacion struct {
	ID                uint       `gorm:"column:hospitalizacion_id;primaryKey" json:"id"`
	PacienteID        uint       `gorm:"column:paciente_id" json:"paciente_id"`
	FechaIngreso      time.Time  `gorm:"column:fecha_ingreso" json:"fecha_ingreso"`
	FechaAlta         *time.Time `gorm:"column:fecha_alta" json:"fecha_alta"`
	Motivo            string     `gorm:"column:motivo" json:"motivo"`
	SalaID            *uint      `gorm:"column:sala_id" json:"sala_id"`
	MedicoResponsable *uint      `gorm:"column:medico_responsable" json:"medico_responsable"`
	Notas             string     `gorm:"column:notas" json:"notas"`
}

func (Hospitalizacion) TableName() string {
	return "Hospitalizacion"
}
