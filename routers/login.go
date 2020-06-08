package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/xfchris/gotter/bd"
	"github.com/xfchris/gotter/jwt"
	"github.com/xfchris/gotter/models"
)

//Login crea token, y lo muestra en pantalla
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Usuario o contraseña incorrecta: "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido ", 400)
		return
	}

	usuario, existe := bd.IntentoLogin(t.Email, t.Password)

	if !existe {
		http.Error(w, "Usuario y/o contraseña incorrecta", 400)
		return
	}

	jwtKey, err := jwt.GenerarJWT(usuario)
	if err != nil {
		http.Error(w, "Error al generar el token "+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	//crearCookie(w, jwtKey)
}

func crearCookie(w http.ResponseWriter, jwtKey string) {
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
