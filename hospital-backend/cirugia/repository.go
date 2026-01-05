package cirugia

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAll() ([]Cirugia, error) {
	var lista []Cirugia
	// si en tu DB la columna se llama diferente, cámbiala aquí
	err := r.db.Order("cirugia_id").Find(&lista).Error
	return lista, err
}

func (r *Repository) GetByID(id uint) (*Cirugia, error) {
	var c Cirugia
	err := r.db.First(&c, id).Error
	return &c, err
}

func (r *Repository) GetByPaciente(pacienteID uint) ([]Cirugia, error) {
	var lista []Cirugia
	err := r.db.Where("paciente_id = ?", pacienteID).Find(&lista).Error
	return lista, err
}

func (r *Repository) Create(c *Cirugia) error {
	return r.db.Create(c).Error
}

func (r *Repository) Update(c *Cirugia) error {
	return r.db.Save(c).Error
}

func (r *Repository) Delete(id uint) error {
	return r.db.Delete(&Cirugia{}, id).Error
}
