package tipo_prueba

type TipoPrueba struct {
	TipoPruebaID int     `gorm:"column:tipo_prueba_id;primaryKey" json:"tipo_prueba_id"`
	Nombre       string  `gorm:"column:nombre;size:100;not null;unique" json:"nombre"`
	Descripcion  *string `gorm:"column:descripcion" json:"descripcion,omitempty"`
}

func (TipoPrueba) TableName() string {
	return "Tipo_Prueba"
}
