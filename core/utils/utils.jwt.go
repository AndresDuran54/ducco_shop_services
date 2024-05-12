package utils

import (
	"github.com/dgrijalva/jwt-go"
)

type UtilsJWT struct{}

func (uj UtilsJWT) GenerateJWTToken(secretKey string, expTimestamp uint64, customerId uint32) (string, error) {
	//+ Creamos el token con un conjunto de reclamos (claims)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	//+ Agregamos reclamos al token
	claims["customerId"] = customerId
	claims["exp"] = expTimestamp

	//+ Firmamos el token con la clave secreta
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
