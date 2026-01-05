package usuario

type Usuario struct {
	UsuarioID          int     `gorm:"column:usuario_id;primaryKey" json:"usuario_id"`
	Username           string  `gorm:"column:username;unique" json:"username"`
	PasswordHash       string  `gorm:"column:password_hash" json:"-"`
	Email              *string `gorm:"column:email" json:"email,omitempty"`
	Activo             bool    `gorm:"column:activo" json:"activo"`
	MedicoIDAsociado   *int    `gorm:"column:medico_id_asociado" json:"medico_id_asociado,omitempty"`
	PacienteIDAsociado *int    `gorm:"column:paciente_id_asociado" json:"paciente_id_asociado,omitempty"`
}

func (Usuario) TableName() string {
	return "Usuario"
}
