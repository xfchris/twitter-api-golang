package bd

import (
	"github.com/xfchris/gotter/models"
	"golang.org/x/crypto/bcrypt"
)

//IntentoLogin comprueba credenciales y devuelve usuario si la comprobacion fue exitosa
func IntentoLogin(email string, password string) (models.Usuario, bool) {

	usuario, encontrado, _ := ExisteEmail(email)

	if encontrado == false {
		return usuario, false
	}

	claveHash := []byte(usuario.Password)
	clave := []byte(password)
	err := bcrypt.CompareHashAndPassword(claveHash, clave)

	if err != nil {
		//var usuarioVacio models.Usuario
		return usuario, false
	}

	return usuario, true
}
