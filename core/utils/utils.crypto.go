package utils

import (
	"golang.org/x/crypto/bcrypt"
)

type UtilsCrypto struct{}

func (uc UtilsCrypto) GenerateFromPassword(password string) (string, error) {
	//+ Generamos el hash de la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	return string(hashedPassword), err
}

func (uc UtilsCrypto) CompareHashAndPassword(hashedPassword string, password string) error {
	// Verificar la contraseña encriptada
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err
}
