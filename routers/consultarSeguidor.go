package routers

import (
	"encoding/json"
	"net/http"

	"github.com/xfchris/gotter/bd"

	"github.com/xfchris/gotter/middle"
	"github.com/xfchris/gotter/models"
)

//ConsultarSeguidor Para saber si yo sigo al usuario enviado en el GET
func ConsultarSeguidor(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.UsuarioSeguido
	t.UsuarioID = middle.IDUsuario
	t.UsuarioSeguido = ID

	var resp = map[string]bool{
		"status": false,
	}
	//resp := "{\"status\":false}"

	status, err := bd.ConsultoSeguidor(t)
	if !(err != nil || !status) {
		//resp = "{\"status\":true}"
		resp["status"] = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
	//fmt.Fprintln(w, resp)
}
