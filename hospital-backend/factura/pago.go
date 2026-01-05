package facturacion

type Pago struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	Metodo     string  `json:"metodo"` // EFECTIVO, TARJETA, TRANSFERENCIA
	Monto      float64 `json:"monto"`
	Referencia string  `json:"referencia"`

	FacturaID uint `json:"factura_id"`
}

func (Pago) TableName() string {
	return "Pago"
}
