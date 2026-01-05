package facturacion

type Aseguradora struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Nombre string `json:"nombre"`
	Ruc    string `json:"ruc"`
}

func (Aseguradora) TableName() string {
	return "Aseguradora"
}
