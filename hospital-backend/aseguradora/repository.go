package aseguradora

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]Aseguradora, error)
	GetByID(id uint) (*Aseguradora, error)
	Create(a *Aseguradora) error
	Update(a *Aseguradora) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Aseguradora, error) {
	var list []Aseguradora
	err := r.db.Find(&list).Error
	return list, err
}

func (r *repository) GetByID(id uint) (*Aseguradora, error) {
	var a Aseguradora
	err := r.db.First(&a, id).Error
	return &a, err
}

func (r *repository) Create(a *Aseguradora) error {
	return r.db.Create(a).Error
}

func (r *repository) Update(a *Aseguradora) error {
	return r.db.Save(a).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Aseguradora{}, id).Error
}
