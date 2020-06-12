package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/xfchris/gotter/bd"
	"github.com/xfchris/gotter/middle"
)

//LeerTweetsSeguidores muestra en json los tweets de seguidores
func LeerTweetsSeguidores(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("page")
	if len(page) < 1 {
		http.Error(w, "Se necesita el parametro page", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "la varaible page debe ser un número", http.StatusBadRequest)
		return
	}

	tweetsSegs, correcto := bd.LeerTweetsSeguidores(middle.IDUsuario, pagina)

	if !correcto {
		http.Error(w, "Error al encontrar los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tweetsSegs)
}
