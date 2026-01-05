package medicamento

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) GetAll() ([]Medicamento, error) {
	var list []Medicamento
	err := r.db.Find(&list).Error
	return list, err
}

func (r *Repository) GetByID(id uint) (*Medicamento, error) {
	var m Medicamento
	err := r.db.First(&m, id).Error
	return &m, err
}

func (r *Repository) Create(m *Medicamento) error {
	return r.db.Create(m).Error
}

func (r *Repository) Update(m *Medicamento) error {
	return r.db.Save(m).Error
}

func (r *Repository) Delete(id uint) error {
	return r.db.Delete(&Medicamento{}, id).Error
}
