package horario

import "time"

type HorarioMedico struct {
	HorarioID  int       `gorm:"column:horario_id;primaryKey;autoIncrement"`
	MedicoID   int       `gorm:"column:medico_id;not null"`
	DiaSemana  string    `gorm:"column:dia_semana;type:varchar(10);not null"`
	HoraInicio time.Time `gorm:"column:hora_inicio;type:time;not null"`
	HoraFin    time.Time `gorm:"column:hora_fin;type:time;not null"`
}

func (HorarioMedico) TableName() string {
	return "Horario_Medico"
}
