package receta

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]Receta, error)
	GetByID(id uint) (*Receta, error)
	Create(r *Receta) error
	Update(r *Receta) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Receta, error) {
	var list []Receta
	err := r.db.
		Preload("Paciente").
		Preload("Medico").
		Preload("Diagnostico").
		Find(&list).Error

	return list, err
}

func (r *repository) GetByID(id uint) (*Receta, error) {
	var rec Receta
	err := r.db.
		Preload("Paciente").
		Preload("Medico").
		Preload("Diagnostico").
		First(&rec, id).Error

	return &rec, err
}

func (r *repository) Create(rec *Receta) error {
	return r.db.Create(rec).Error
}

func (r *repository) Update(rec *Receta) error {
	return r.db.Save(rec).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Receta{}, id).Error
}
