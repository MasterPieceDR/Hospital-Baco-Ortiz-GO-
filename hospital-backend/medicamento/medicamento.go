package medicamento

type Medicamento struct {
	ID           uint    `gorm:"column:medicamento_id;primaryKey" json:"id"`
	Nombre       string  `gorm:"column:nombre" json:"nombre"`
	Descripcion  string  `gorm:"column:descripcion" json:"descripcion"`
	Presentacion string  `gorm:"column:presentacion" json:"presentacion"`
	Precio       float64 `gorm:"column:precio" json:"precio"`
}

func (Medicamento) TableName() string {
	return "Medicamento"
}
