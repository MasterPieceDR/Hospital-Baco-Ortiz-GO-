package tipo_prueba

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAll() ([]TipoPrueba, error) {
	var tipos []TipoPrueba
	if err := r.db.Order("tipo_prueba_id").Find(&tipos).Error; err != nil {
		return nil, err
	}
	return tipos, nil
}

func (r *Repository) GetByID(id int) (*TipoPrueba, error) {
	var t TipoPrueba
	if err := r.db.First(&t, id).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *Repository) Create(t *TipoPrueba) error {
	return r.db.Create(t).Error
}

func (r *Repository) Update(t *TipoPrueba) error {
	return r.db.Save(t).Error
}

func (r *Repository) Delete(id int) error {
	return r.db.Delete(&TipoPrueba{}, id).Error
}
