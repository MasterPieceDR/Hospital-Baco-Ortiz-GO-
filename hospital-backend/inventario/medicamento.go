package inventario

type Medicamento struct {
	MedicamentoID uint `gorm:"primaryKey;column:medicamento_id"`
	Nombre        string
	Descripcion   string
	Proveedor     string
	Categoria     string
}

func (Medicamento) TableName() string {
	return "Medicamento"
}
