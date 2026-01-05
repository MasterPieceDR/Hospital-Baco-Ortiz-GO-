package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"hospital-backend/internal/config"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func NewDB(cfg *config.Config) *gorm.DB {
	// añadimos encrypt=disable para entornos locales; ajusta según tu servidor
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBServer,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error conectando a la base de datos: %v", err)
	}

	// obtener *sql.DB para ping y configuración del pool
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("warning: no se pudo obtener sql.DB: %v", err)
		return db
	}

	// intentar hacer ping, no fatal para permitir pruebas locales
	if err := sqlDB.Ping(); err != nil {
		log.Printf("warning: fallo ping DB: %v", err)
	}

	// configurar pool (ajusta según carga)
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(1 * time.Hour)

	// Crear índices útiles si no existen (mejora rendimiento de consultas grandes)
	// Nota: en SQL Server usamos chequeo en sys.indexes para no intentar recrear índice si ya existe.
	// Estos execs no son fatales; solo logueamos errores.
	if err := db.Exec(`
		IF NOT EXISTS (
			SELECT name FROM sys.indexes WHERE name = 'idx_consulta_fecha_consulta'
		)
		CREATE INDEX idx_consulta_fecha_consulta ON Consulta(fecha_consulta)
	`).Error; err != nil {
		log.Printf("warning: no se pudo crear idx_consulta_fecha_consulta: %v", err)
	}

	if err := db.Exec(`
		IF NOT EXISTS (
			SELECT name FROM sys.indexes WHERE name = 'idx_diagnostico_consulta_id'
		)
		CREATE INDEX idx_diagnostico_consulta_id ON Diagnostico(consulta_id)
	`).Error; err != nil {
		log.Printf("warning: no se pudo crear idx_diagnostico_consulta_id: %v", err)
	}

	// prevenir import no usado
	_ = sql.DB(*new(sql.DB))

	return db
}
