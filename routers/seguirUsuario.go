package routers

import (
	"net/http"

	"github.com/xfchris/gotter/bd"
	"github.com/xfchris/gotter/middle"
	"github.com/xfchris/gotter/models"
)

//SeguirUsuario crea una relacion pasandole el id el usuario que se va a segir ṕor get
func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parametro ID es obligatorio", 400)
		return
	}

	var t models.UsuarioSeguido

	t.UsuarioID = middle.IDUsuario
	t.UsuarioSeguido = ID

	status, err := bd.SeguirUsuario(t)
	if err != nil {
		http.Error(w, "Error al crear la relacion: "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "Error al crear la relacion", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
