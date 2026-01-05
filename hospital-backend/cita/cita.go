package cita

import "time"

type Cita struct {
	ID         uint      `gorm:"column:cita_id;primaryKey" json:"id"`
	PacienteID uint      `gorm:"column:paciente_id" json:"paciente_id"`
	MedicoID   uint      `gorm:"column:medico_id" json:"medico_id"`
	SalaID     *uint     `gorm:"column:sala_id" json:"sala_id,omitempty"`
	FechaHora  time.Time `gorm:"column:fecha_hora" json:"fecha_hora"`
	Estado     string    `gorm:"column:estado" json:"estado"`
	Motivo     string    `gorm:"column:motivo" json:"motivo"`
	Notas      string    `gorm:"column:notas" json:"notas"`
}

func (Cita) TableName() string {
	return "Cita"
}
