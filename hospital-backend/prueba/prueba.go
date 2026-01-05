package prueba

import (
	"time"

	"hospital-backend/tipo_prueba"
)

type Prueba struct {
	PruebaID        uint                   `gorm:"column:prueba_id;primaryKey" json:"prueba_id"`
	PacienteID      uint                   `gorm:"column:paciente_id;not null" json:"paciente_id"`
	ConsultaID      *uint                  `gorm:"column:consulta_id" json:"consulta_id,omitempty"`
	TipoPruebaID    uint                   `gorm:"column:tipo_prueba_id;not null" json:"tipo_prueba_id"`
	FechaSolicitud  time.Time              `gorm:"column:fecha_solicitud;not null" json:"fecha_solicitud"`
	FechaResultado  *time.Time             `gorm:"column:fecha_resultado" json:"fecha_resultado,omitempty"`
	Resultado       *string                `gorm:"column:resultado" json:"resultado,omitempty"`
	Unidad          *string                `gorm:"column:unidad" json:"unidad,omitempty"`
	ValorReferencia *string                `gorm:"column:valor_referencia" json:"valor_referencia,omitempty"`
	TipoPrueba      tipo_prueba.TipoPrueba `gorm:"foreignKey:TipoPruebaID;references:TipoPruebaID" json:"tipo_prueba"`
}

func (Prueba) TableName() string {
	return "Prueba"
}
