package main

import (
	"log"
	"strings"
	"time"

	"hospital-backend/internal/config"
	"hospital-backend/internal/database"
	"hospital-backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// intentar cargar .env (no es fatal si no existe; en producción se usan variables de entorno)
	if err := godotenv.Load(); err != nil {
		log.Println(".env no cargado, usando variables de entorno del sistema")
	}

	cfg := config.LoadConfig()
	db := database.NewDB(cfg)

	r := gin.Default()

	// configurar AllowOrigins desde FRONTEND_URL (coma-separado) o "*" por defecto
	var origins []string
	if cfg.FrontendURL == "*" || cfg.FrontendURL == "" {
		origins = []string{"*"}
	} else {
		origins = strings.Split(cfg.FrontendURL, ",")
	}

	// CORS: permitir llamadas desde el frontend (ajusta AllowOrigins según tu dominio en producción)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     origins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// endpoint simple de health check público
	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	routes.Register(r, db)

	addr := ":" + cfg.AppPort
	log.Println("Servidor escuchando en", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
