package medico

type Medico struct {
	ID              uint   `gorm:"primaryKey;column:medico_id" json:"id"`
	Nombre          string `gorm:"column:nombre" json:"nombre"`
	Apellido        string `gorm:"column:apellidos" json:"apellido"`
	Correo          string `gorm:"column:email" json:"correo"`
	Telefono        string `gorm:"column:telefono" json:"telefono"`
	NumeroColegiado string `gorm:"column:numero_colegiado" json:"numero_colegiado"`
	Activo          bool   `gorm:"column:activo" json:"activo"`
}

func (Medico) TableName() string {
	return "Medico"
}
