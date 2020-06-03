package routers

import (
	"encoding/json"
	"net/http"

	"github.com/xfchris/gotter/bd"
	"github.com/xfchris/gotter/models"
)

//Registro permite crear un registro en base de datos
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos del registro: "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email esta vacio", 400)
	}

	if len(t.Password) < 6 {
		http.Error(w, "Password debe ser de 6 caracteres", 400)
	}

	_, encontrado, _ := bd.ExisteEmail(t.Email)
	if encontrado {
		http.Error(w, "Ya existe ese email", 400)
		return
	}

	_, status, err := bd.CrearUsuario(t)
	if err != nil {
		http.Error(w, "Error al crear usuario "+err.Error(), 400)
		return
	}
	//por si no guardo en mongo a pesar que no genero error
	if status == false {
		http.Error(w, "No se ha podido crear el usuario.", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
