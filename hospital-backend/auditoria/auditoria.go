package auditoria

import "time"

type Auditoria struct {
	AuditoriaID int       `gorm:"column:auditoria_id;primaryKey;autoIncrement" json:"auditoria_id"`
	UsuarioID   *int      `gorm:"column:usuario_id" json:"usuario_id"`
	FechaHora   time.Time `gorm:"column:fecha_hora" json:"fecha_hora"`
	Tabla       string    `gorm:"column:tabla" json:"tabla"`
	RegistroID  *int      `gorm:"column:registro_id" json:"registro_id"`
	Accion      string    `gorm:"column:accion" json:"accion"`
	Descripcion string    `gorm:"column:descripcion" json:"descripcion"`
}

func (Auditoria) TableName() string {
	return "Auditoria"
}
