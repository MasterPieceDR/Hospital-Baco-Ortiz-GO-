package rol

type Rol struct {
	RolID       uint    `gorm:"column:rol_id;primaryKey" json:"rol_id"`
	Nombre      string  `gorm:"column:nombre;size:50;not null" json:"nombre"`
	Descripcion *string `gorm:"column:descripcion" json:"descripcion,omitempty"`
}

func (Rol) TableName() string {
	return "Rol"
}
