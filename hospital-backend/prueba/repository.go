package prueba

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAll() ([]Prueba, error) {
	var pruebas []Prueba
	if err := r.db.Preload("TipoPrueba").Order("prueba_id").Find(&pruebas).Error; err != nil {
		return nil, err
	}
	return pruebas, nil
}

func (r *Repository) GetByID(id uint) (*Prueba, error) {
	var p Prueba
	if err := r.db.Preload("TipoPrueba").First(&p, id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *Repository) GetByPacienteID(pacienteID uint) ([]Prueba, error) {
	var pruebas []Prueba
	if err := r.db.Preload("TipoPrueba").
		Where("paciente_id = ?", pacienteID).
		Order("fecha_solicitud DESC").
		Find(&pruebas).Error; err != nil {
		return nil, err
	}
	return pruebas, nil
}

func (r *Repository) Create(p *Prueba) error {
	return r.db.Create(p).Error
}

func (r *Repository) Update(p *Prueba) error {
	return r.db.Save(p).Error
}

func (r *Repository) Delete(id uint) error {
	return r.db.Delete(&Prueba{}, id).Error
}
