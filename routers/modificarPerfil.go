package routers

import (
	"encoding/json"
	"net/http"

	"github.com/xfchris/gotter/bd"
	"github.com/xfchris/gotter/middle"
	"github.com/xfchris/gotter/models"
)

//ModificarPerfil actualiza un perfil de usuario
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	//obtengo datos
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos: "+err.Error(), 400)
		return
	}
	editado, err := bd.ModificarRegistro(t, middle.IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un error al editar usuario: "+err.Error(), 400)
		return
	}
	if !editado {
		http.Error(w, "No se pudo editar el usuario: ", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
