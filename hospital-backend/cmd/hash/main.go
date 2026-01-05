package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	passwords := map[string]string{
		"admin":     "admin123",
		"ltorres":   "medico123",
		"csanchez":  "medico456",
		"scastro":   "recepcion123",
		"agarcia":   "paciente123",
		"lmartinez": "paciente456",
		"rmena":     "enfermeria123",
		"pluna":     "finanzas123",
		"mparedes":  "farmacia123",
		"dsantos":   "director123",
	}

	for user, pass := range passwords {
		hash, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
		fmt.Println(user, "=>", string(hash))
	}
}
