package config

import "os"

type Config struct {
	DBServer    string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
	JWTSecret   string
	AppPort     string
	FrontendURL string
}

func LoadConfig() *Config {
	return &Config{
		DBServer:    getEnv("DB_SERVER", "localhost"),
		DBPort:      getEnv("DB_PORT", "1433"),
		DBUser:      getEnv("DB_USER", "goproject"),
		DBPassword:  getEnv("DB_PASSWORD", "root"),
		DBName:      getEnv("DB_NAME", "HospitalDB"),
		JWTSecret:   getEnv("JWT_SECRET", "supersecretjwt"),
		AppPort:     getEnv("APP_PORT", "8080"),
		FrontendURL: getEnv("FRONTEND_URL", "*"),
	}
}

func getEnv(key, defaultValue string) string {
	v := os.Getenv(key)
	if v == "" {
		return defaultValue
	}
	return v
}
