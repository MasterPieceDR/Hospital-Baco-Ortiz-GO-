package consulta

import (
	"time"

	"hospital-backend/diagnostico"
)

type Consulta struct {
	ID         uint      `gorm:"column:consulta_id;primaryKey" json:"id"`
	CitaID     *uint     `gorm:"column:cita_id" json:"cita_id,omitempty"`
	MedicoID   uint      `gorm:"column:medico_id" json:"medico_id"`
	PacienteID uint      `gorm:"column:paciente_id" json:"paciente_id"`
	Fecha      time.Time `gorm:"column:fecha_consulta" json:"fecha"`
	Motivo     string    `gorm:"column:motivo" json:"motivo"`
	Notas      string    `gorm:"column:notas_medicas" json:"notas_medicas"`
	// Diagnósticos relacionados (preload desde repo)
	Diagnosticos []diagnostico.Diagnostico `gorm:"foreignKey:ConsultaID;references:ID" json:"diagnosticos,omitempty"`
}

func (Consulta) TableName() string {
	return "Consulta"
}
