package routers

import (
	"net/http"

	"github.com/xfchris/gotter/bd"
	"github.com/xfchris/gotter/middle"
)

//BorrarTweet borra el tweet
func BorrarTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Se necesita el id", 400)
		return
	}
	err := bd.BorrarTweet(ID, middle.IDUsuario)
	if err != nil {
		http.Error(w, "Error al eliminar tweet "+err.Error(), 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
