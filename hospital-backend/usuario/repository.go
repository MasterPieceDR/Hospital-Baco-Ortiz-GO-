package usuario

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindByUsername(username string) (*Usuario, error) {
	var u Usuario
	err := r.db.Where("username = ?", username).First(&u).Error
	return &u, err
}

func (r *Repository) GetRolesByUsuarioID(id int) ([]string, error) {
	var roles []string
	err := r.db.Raw(`
		SELECT R.nombre
		FROM Rol R
		INNER JOIN Usuario_Rol UR ON UR.rol_id = R.rol_id
		WHERE UR.usuario_id = ?
	`, id).Scan(&roles).Error
	return roles, err
}
