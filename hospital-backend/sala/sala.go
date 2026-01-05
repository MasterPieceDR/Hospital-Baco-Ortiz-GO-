package sala

type Sala struct {
	ID          uint   `gorm:"column:sala_id;primaryKey" json:"id"`
	Nombre      string `gorm:"column:nombre" json:"nombre"`
	Tipo        string `gorm:"column:tipo" json:"tipo"`
	Ubicacion   string `gorm:"column:ubicacion" json:"ubicacion"`
	Descripcion string `gorm:"column:descripcion" json:"descripcion"`
}

func (Sala) TableName() string {
	return "Sala"
}
