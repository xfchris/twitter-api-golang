package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/xfchris/gotter/bd"
	jwtLocal "github.com/xfchris/gotter/jwt"
	"github.com/xfchris/gotter/models"
)

//Email contiene el email del que inicio sesion
var Email string

//IDUsuario contiene el id del usuario que inició sesion
var IDUsuario string

//ProcesarToken procesa el token y nos dice si es correcto
func ProcesarToken(aHash string) (*models.Claim, bool, string, error) {

	claims := &models.Claim{}
	splitToken := strings.Split(aHash, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, "", errors.New("Formato de token invalido")
	}
	aHash = strings.TrimSpace(splitToken[1])
	//Valida el token y si lo valida llena la estructura clain
	tkn, err := jwt.ParseWithClaims(aHash, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtLocal.ClaveSecreta, nil
	})
	if err == nil {
		_, encontrado, _ := bd.ExisteEmail(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, "", errors.New("Token inválido")
	}
	return claims, false, "", err
}
