package facturacion

type Poliza struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	Numero        string `gorm:"unique" json:"numero"`
	Descripcion   string `json:"descripcion"`
	Cobertura     string `json:"cobertura"` // total/ parcial
	AseguradoraID uint   `json:"aseguradora_id"`
}

func (Poliza) TableName() string {
	return "Poliza"
}
