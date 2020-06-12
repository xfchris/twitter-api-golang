package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/xfchris/gotter/bd"
	"github.com/xfchris/gotter/middle"
)

//LeerUsuarios muestra los usuarios pasados por parametros get
func LeerUsuarios(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagInt, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Debe enviar el parametro pagina", 400)
		return
	}

	results, status := bd.LeerUsuarios(middle.IDUsuario, pagInt, search, typeUser)
	if !status {
		http.Error(w, "Error al leer usuarios", 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&results)
}
