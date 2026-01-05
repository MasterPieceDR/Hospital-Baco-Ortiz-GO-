package usuario_rol

type UsuarioRol struct {
	UsuarioID int `gorm:"column:usuario_id;primaryKey" json:"usuario_id"`
	RolID     int `gorm:"column:rol_id;primaryKey" json:"rol_id"`
}

func (UsuarioRol) TableName() string {
	return "Usuario_Rol"
}
