package receta

import (
	diagnostico "hospital-backend/diagnostico"
	medico "hospital-backend/medico"
	paciente "hospital-backend/paciente"
	"time"
)

type Receta struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	PacienteID    uint      `json:"paciente_id"`
	MedicoID      uint      `json:"medico_id"`
	DiagnosticoID uint      `json:"diagnostico_id"`
	Fecha         time.Time `json:"fecha"`
	Indicaciones  string    `json:"indicaciones"`

	Paciente    paciente.Paciente       `gorm:"foreignKey:PacienteID" json:"paciente"`
	Medico      medico.Medico           `gorm:"foreignKey:MedicoID" json:"medico"`
	Diagnostico diagnostico.Diagnostico `gorm:"foreignKey:DiagnosticoID" json:"diagnostico"`
}

func (Receta) TableName() string {
	return "Receta"
}
