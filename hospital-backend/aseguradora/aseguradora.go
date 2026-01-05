package aseguradora

type Aseguradora struct {
	ID        uint   `gorm:"primaryKey;column:aseguradora_id" json:"id"`
	Nombre    string `gorm:"column:nombre" json:"nombre"`
	Telefono  string `gorm:"column:telefono" json:"telefono"`
	Direccion string `gorm:"column:direccion" json:"direccion"`
}

func (Aseguradora) TableName() string {
	return "Aseguradora"
}
