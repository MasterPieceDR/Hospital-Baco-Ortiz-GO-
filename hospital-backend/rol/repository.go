package rol

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAll() ([]Rol, error) {
	var roles []Rol
	if err := r.db.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *Repository) GetByID(id uint) (*Rol, error) {
	var role Rol
	if err := r.db.First(&role, "rol_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *Repository) Create(role *Rol) error {
	return r.db.Create(role).Error
}

func (r *Repository) Update(role *Rol) error {
	return r.db.Save(role).Error
}

func (r *Repository) Delete(id uint) error {
	return r.db.Delete(&Rol{}, "rol_id = ?", id).Error
}
