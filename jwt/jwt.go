package jwt

import(
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/xfchris/gotter/models"
)
//ClaveSecreta para el jwt
var ClaveSecreta = []byte("ClaveSecreta_Chris")

//GenerarJWT genera un token JWT y devuelve una cadena JWT firmada
func GenerarJWT(t models.Usuario) (string, error){

	payload := jwt.MapClaims{
		"email": t.Email,
		"nombres": t.Nombres,
		"apellidos": t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia": t.Biografia,
		"ubicacion": t.Ubicacion,
		"sitioWeb": t.SitioWeb,
		"_id": t.ID.Hex(),
		"exp": time.Now().Add(24*time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(ClaveSecreta)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}