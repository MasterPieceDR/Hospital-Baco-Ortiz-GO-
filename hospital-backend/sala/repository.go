package sala

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) GetAll() ([]Sala, error) {
	var salas []Sala
	err := r.db.Find(&salas).Error
	return salas, err
}

func (r *Repository) GetByID(id uint) (*Sala, error) {
	var sala Sala
	err := r.db.First(&sala, id).Error
	return &sala, err
}

func (r *Repository) Create(s *Sala) error {
	return r.db.Create(s).Error
}

func (r *Repository) Update(s *Sala) error {
	return r.db.Save(s).Error
}

func (r *Repository) Delete(id uint) error {
	return r.db.Delete(&Sala{}, id).Error
}
