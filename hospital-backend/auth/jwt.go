package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserID     int      `json:"user_id"`
	Username   string   `json:"username"`
	Roles      []string `json:"roles"`
	MedicoID   *int     `json:"medico_id,omitempty"`
	PacienteID *int     `json:"paciente_id,omitempty"`
	jwt.RegisteredClaims
}

func getSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "supersecretjwt"
	}
	return []byte(secret)
}

func GenerateJWT(userID int, username string, roles []string, medicoID, pacienteID *int) (string, error) {
	now := time.Now()

	claims := CustomClaims{
		UserID:     userID,
		Username:   username,
		Roles:      roles,
		MedicoID:   medicoID,
		PacienteID: pacienteID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "hospital-backend",
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getSecret())
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return getSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}
	return claims, nil
}
