package usuario_rol

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetRolesByUsuarioID(usuarioID int) ([]int, error) {
	var userRoles []UsuarioRol
	if err := r.db.Where("usuario_id = ?", usuarioID).Find(&userRoles).Error; err != nil {
		return nil, err
	}
	roles := make([]int, len(userRoles))
	for i, ur := range userRoles {
		roles[i] = ur.RolID
	}
	return roles, nil
}

func (r *Repository) AddRole(usuarioID, rolID int) error {
	record := UsuarioRol{
		UsuarioID: usuarioID,
		RolID:     rolID,
	}
	return r.db.Create(&record).Error
}

func (r *Repository) RemoveRole(usuarioID, rolID int) error {
	return r.db.Where("usuario_id = ? AND rol_id = ?", usuarioID, rolID).Delete(&UsuarioRol{}).Error
}
