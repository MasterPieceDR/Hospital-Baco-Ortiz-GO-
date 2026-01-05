package hospitalizacion

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) GetAll() ([]Hospitalizacion, error) {
	var list []Hospitalizacion
	err := r.db.Find(&list).Error
	return list, err
}

func (r *Repository) GetByID(id uint) (*Hospitalizacion, error) {
	var h Hospitalizacion
	err := r.db.First(&h, id).Error
	return &h, err
}

func (r *Repository) Create(h *Hospitalizacion) error {
	return r.db.Create(h).Error
}

func (r *Repository) Update(h *Hospitalizacion) error {
	return r.db.Save(h).Error
}

func (r *Repository) Delete(id uint) error {
	return r.db.Delete(&Hospitalizacion{}, id).Error
}
