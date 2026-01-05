package personal

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) GetAll() ([]Personal, error) {
	var list []Personal
	err := r.db.Find(&list).Error
	return list, err
}

func (r *Repository) GetByID(id uint) (*Personal, error) {
	var p Personal
	err := r.db.First(&p, id).Error
	return &p, err
}

func (r *Repository) Create(p *Personal) error {
	return r.db.Create(p).Error
}

func (r *Repository) Update(p *Personal) error {
	return r.db.Save(p).Error
}

func (r *Repository) Delete(id uint) error {
	return r.db.Delete(&Personal{}, id).Error
}
