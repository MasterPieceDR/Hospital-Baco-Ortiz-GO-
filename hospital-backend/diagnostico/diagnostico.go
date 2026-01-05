package diagnostico

import "time"

type Diagnostico struct {
	DiagnosticoID uint      `gorm:"column:diagnostico_id;primaryKey" json:"diagnostico_id"`
	ConsultaID    uint      `gorm:"column:consulta_id;not null" json:"consulta_id"`
	Descripcion   string    `gorm:"column:descripcion;not null" json:"descripcion"`
	Fecha         time.Time `gorm:"column:fecha;not null" json:"fecha"`
}

func (Diagnostico) TableName() string {
	return "Diagnostico"
}
