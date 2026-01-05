package pago

import "time"

type Pago struct {
	ID            uint      `gorm:"primaryKey;column:pago_id" json:"id"`
	Monto         float64   `gorm:"column:monto" json:"monto"`
	FechaPago     time.Time `gorm:"column:fecha_pago" json:"fecha_pago"`
	MetodoPago    string    `gorm:"column:metodo_pago" json:"metodo_pago"`
	FacturaID     uint      `gorm:"column:factura_id" json:"factura_id"`
	Observaciones string    `gorm:"column:observaciones" json:"observaciones"`
}

func (Pago) TableName() string {
	return "pago"
}
