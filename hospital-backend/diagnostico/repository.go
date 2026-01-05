package diagnostico

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAll() ([]Diagnostico, error) {
	var lista []Diagnostico
	if err := r.db.Order("diagnostico_id").Find(&lista).Error; err != nil {
		return nil, err
	}
	return lista, nil
}

func (r *Repository) GetByID(id uint) (*Diagnostico, error) {
	var d Diagnostico
	if err := r.db.First(&d, id).Error; err != nil {
		return nil, err
	}
	return &d, nil
}

func (r *Repository) GetByConsulta(consultaID uint) ([]Diagnostico, error) {
	var lista []Diagnostico
	if err := r.db.Where("consulta_id = ?", consultaID).Find(&lista).Error; err != nil {
		return nil, err
	}
	return lista, nil
}

func (r *Repository) Create(data *Diagnostico) error {
	return r.db.Create(data).Error
}

func (r *Repository) Update(data *Diagnostico) error {
	return r.db.Save(data).Error
}

func (r *Repository) Delete(id uint) error {
	return r.db.Delete(&Diagnostico{}, id).Error
}
