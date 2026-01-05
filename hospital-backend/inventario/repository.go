package inventario

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]Inventario, error)
	GetByID(id uint) (*Inventario, error)
	Create(inv *Inventario) error
	Update(inv *Inventario) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Inventario, error) {
	var inv []Inventario
	err := r.db.Preload("Medicamento").Find(&inv).Error
	return inv, err
}

func (r *repository) GetByID(id uint) (*Inventario, error) {
	var inv Inventario
	err := r.db.Preload("Medicamento").First(&inv, id).Error
	return &inv, err
}

func (r *repository) Create(inv *Inventario) error {
	return r.db.Create(inv).Error
}

func (r *repository) Update(inv *Inventario) error {
	return r.db.Save(inv).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Inventario{}, id).Error
}
