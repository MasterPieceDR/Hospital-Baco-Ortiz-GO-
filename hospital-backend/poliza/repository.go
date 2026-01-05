package poliza

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]Poliza, error)
	GetByID(id uint) (*Poliza, error)
	Create(p *Poliza) error
	Update(p *Poliza) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Poliza, error) {
	var list []Poliza
	err := r.db.Preload("Paciente").Preload("Aseguradora").Find(&list).Error
	return list, err
}

func (r *repository) GetByID(id uint) (*Poliza, error) {
	var p Poliza
	err := r.db.Preload("Paciente").Preload("Aseguradora").
		First(&p, id).Error
	return &p, err
}

func (r *repository) Create(p *Poliza) error {
	return r.db.Create(p).Error
}

func (r *repository) Update(p *Poliza) error {
	return r.db.Save(p).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Poliza{}, id).Error
}
