package inventario

import (
	"hospital-backend/medicamento"
	"time"
)

type Inventario struct {
	ID            uint                    `gorm:"primaryKey" json:"id"`
	MedicamentoID uint                    `json:"medicamento_id"`
	Medicamento   medicamento.Medicamento `gorm:"foreignKey:MedicamentoID" json:"medicamento"`
	Cantidad      int                     `json:"cantidad"`
	FechaIngreso  time.Time               `json:"fecha_ingreso"`
}

func (Inventario) TableName() string {
	return "inventarios"
}
