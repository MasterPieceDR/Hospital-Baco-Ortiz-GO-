package personal

type Personal struct {
	ID           uint   `gorm:"column:personal_id;primaryKey" json:"id"`
	Nombre       string `gorm:"column:nombre" json:"nombre"`
	Apellidos    string `gorm:"column:apellidos" json:"apellidos"`
	Rol          string `gorm:"column:rol" json:"rol"`
	Departamento string `gorm:"column:departamento" json:"departamento"`
	Telefono     string `gorm:"column:telefono" json:"telefono"`
	Email        string `gorm:"column:email" json:"email"`
}

func (Personal) TableName() string {
	return "Personal"
}
