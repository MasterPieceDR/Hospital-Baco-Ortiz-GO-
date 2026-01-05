package pago

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindAll() ([]Pago, error) {
	var p []Pago
	err := r.db.Find(&p).Error
	return p, err
}

func (r *Repository) FindByID(id uint) (*Pago, error) {
	var p Pago
	err := r.db.First(&p, id).Error
	return &p, err
}

func (r *Repository) Create(p *Pago) error {
	return r.db.Create(p).Error
}

func (r *Repository) Update(p *Pago) error {
	return r.db.Save(p).Error
}

func (r *Repository) Delete(id uint) error {
	return r.db.Delete(&Pago{}, id).Error
}
