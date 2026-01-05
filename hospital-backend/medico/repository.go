package medico

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) FindAll() ([]Medico, error) {
	var medicos []Medico
	err := r.db.Find(&medicos).Error
	return medicos, err
}

func (r *Repository) FindByID(id uint) (*Medico, error) {
	var m Medico
	err := r.db.First(&m, id).Error
	return &m, err
}

func (r *Repository) Create(m *Medico) error {
	return r.db.Create(m).Error
}

func (r *Repository) Update(m *Medico) error {
	return r.db.Save(m).Error
}

func (r *Repository) Delete(id uint) error {
	return r.db.Delete(&Medico{}, id).Error
}
