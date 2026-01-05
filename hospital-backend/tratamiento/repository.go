package tratamiento

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]Tratamiento, error)
	GetByID(id int) (*Tratamiento, error)
	Create(t *Tratamiento) error
	Update(t *Tratamiento) error
	Delete(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll() ([]Tratamiento, error) {
	var list []Tratamiento
	err := r.db.Find(&list).Error
	return list, err
}

func (r *repository) GetByID(id int) (*Tratamiento, error) {
	var t Tratamiento
	err := r.db.First(&t, id).Error
	return &t, err
}

func (r *repository) Create(t *Tratamiento) error {
	return r.db.Create(t).Error
}

func (r *repository) Update(t *Tratamiento) error {
	return r.db.Save(t).Error
}

func (r *repository) Delete(id int) error {
	return r.db.Delete(&Tratamiento{}, id).Error
}
