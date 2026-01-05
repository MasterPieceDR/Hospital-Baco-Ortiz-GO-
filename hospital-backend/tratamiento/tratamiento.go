package tratamiento

type Tratamiento struct {
	TratamientoID int    `gorm:"column:tratamiento_id;primaryKey" json:"tratamiento_id"`
	ConsultaID    int    `gorm:"column:consulta_id;not null" json:"consulta_id"`
	Descripcion   string `gorm:"column:descripcion;not null" json:"descripcion"`
	DuracionDias  *int   `gorm:"column:duracion_dias" json:"duracion_dias,omitempty"`
}

func (Tratamiento) TableName() string {
	return "Tratamiento"
}
