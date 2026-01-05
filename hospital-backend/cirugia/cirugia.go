package cirugia

import "time"

type Cirugia struct {
	CirugiaID         uint      `gorm:"column:cirugia_id;primaryKey;autoIncrement" json:"cirugia_id"`
	HospitalizacionID uint      `gorm:"column:hospitalizacion_id" json:"hospitalizacion_id"`
	PacienteID        uint      `gorm:"column:paciente_id;not null" json:"paciente_id"`
	MedicoID          uint      `gorm:"column:medico_id;not null" json:"medico_id"`
	AnestesistaID     *uint     `gorm:"column:anestesista_id" json:"anestesista_id,omitempty"`
	FechaCirugia      time.Time `gorm:"column:fecha_cirugia;not null" json:"fecha_cirugia"`
	SalaID            uint      `gorm:"column:sala_id" json:"sala_id"`
	Tipo              string    `gorm:"column:tipo" json:"tipo"`
	Notas             string    `gorm:"column:notas" json:"notas"`
}

func (Cirugia) TableName() string {
	// nombre exacto de la tabla en SQL Server
	return "Cirugia"
}
