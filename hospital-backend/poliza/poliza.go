package poliza

import (
	"hospital-backend/aseguradora"
	"hospital-backend/paciente"
	"time"
)

type Poliza struct {
	ID            uint      `gorm:"primaryKey;column:poliza_id" json:"id"`
	PacienteID    uint      `gorm:"column:paciente_id" json:"paciente_id"`
	AseguradoraID uint      `gorm:"column:aseguradora_id" json:"aseguradora_id"`
	NumeroPoliza  string    `gorm:"column:numero_poliza" json:"numero_poliza"`
	FechaInicio   time.Time `gorm:"column:fecha_inicio" json:"fecha_inicio"`
	FechaFin      time.Time `gorm:"column:fecha_fin" json:"fecha_fin"`
	Cobertura     string    `gorm:"column:cobertura" json:"cobertura"`

	Paciente    paciente.Paciente       `gorm:"foreignKey:PacienteID" json:"paciente"`
	Aseguradora aseguradora.Aseguradora `gorm:"foreignKey:AseguradoraID" json:"aseguradora"`
}

func (Poliza) TableName() string {
	return "Poliza"
}
