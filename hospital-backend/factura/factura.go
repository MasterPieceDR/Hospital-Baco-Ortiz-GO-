package facturacion

import (
	"time"
)

type Factura struct {
	ID     uint      `gorm:"primaryKey" json:"id"`
	Numero string    `gorm:"unique" json:"numero"`
	Fecha  time.Time `json:"fecha"`
	Total  float64   `json:"total"`
	Estado string    `json:"estado"` // PAGADA, PENDIENTE, ANULADA

	PacienteID uint        `json:"paciente_id"`
	Paciente   interface{} `gorm:"-" json:"paciente,omitempty"` // Se carga en el handler si deseas

	PagoID uint `json:"pago_id"`
	Pago   Pago `json:"pago"`

	PolizaID uint   `json:"poliza_id"`
	Poliza   Poliza `json:"poliza"`

	AseguradoraID uint        `json:"aseguradora_id"`
	Aseguradora   Aseguradora `json:"aseguradora"`
}

func (Factura) TableName() string {
	return "Factura"
}
