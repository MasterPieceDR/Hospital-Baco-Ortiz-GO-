package receta

type RecetaMedicamento struct {
	RecetaID      int    `gorm:"column:receta_id;primaryKey"`
	MedicamentoID int    `gorm:"column:medicamento_id;primaryKey"`
	Dosis         string `gorm:"column:dosis;type:varchar(50);not null"`
	Frecuencia    string `gorm:"column:frecuencia;type:varchar(50);not null"`
	DuracionDias  int    `gorm:"column:duracion_dias"`
	Unidad        string `gorm:"column:unidad;type:varchar(20)"`
}

func (RecetaMedicamento) TableName() string {
	return "Receta_Medicamento"
}
