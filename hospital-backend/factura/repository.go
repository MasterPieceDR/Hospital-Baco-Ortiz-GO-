package facturacion

import "gorm.io/gorm"

type Repository interface {
	CreateFactura(f Factura) error
	GetFactura(id uint) (*Factura, error)
	GetAllFacturas() ([]Factura, error)
	UpdateFactura(f Factura) error
	DeleteFactura(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) CreateFactura(f Factura) error {
	return r.db.Create(&f).Error
}

func (r *repository) GetFactura(id uint) (*Factura, error) {
	var f Factura
	err := r.db.Preload("Pago").Preload("Poliza").Preload("Aseguradora").
		First(&f, id).Error
	return &f, err
}

func (r *repository) GetAllFacturas() ([]Factura, error) {
	var facturas []Factura
	err := r.db.Preload("Pago").Preload("Poliza").Preload("Aseguradora").
		Find(&facturas).Error
	return facturas, err
}

func (r *repository) UpdateFactura(f Factura) error {
	return r.db.Save(&f).Error
}

func (r *repository) DeleteFactura(id uint) error {
	return r.db.Delete(&Factura{}, id).Error
}
