package paciente

import "time"

type Paciente struct {
	ID              uint      `gorm:"primaryKey;column:paciente_id" json:"id"`
	Nombre          string    `gorm:"column:nombre" json:"nombre"`
	Apellido        string    `gorm:"column:apellidos" json:"apellido"`
	Cedula          string    `gorm:"column:numero_seguridad_social;unique" json:"cedula"`
	Correo          string    `gorm:"column:email;unique" json:"correo"`
	Telefono        string    `gorm:"column:telefono" json:"telefono"`
	Direccion       string    `gorm:"column:direccion" json:"direccion"`
	FechaNacimiento time.Time `gorm:"column:fecha_nacimiento" json:"fecha_nacimiento"`
	Estado          bool      `gorm:"column:estado" json:"estado"`
	FechaCreacion   time.Time `gorm:"column:fecha_creacion" json:"fecha_creacion"`
}

func (Paciente) TableName() string {
	return "Paciente"
}
