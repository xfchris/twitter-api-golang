package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/xfchris/gotter/bd"
)

//LeerTweets te lista en pantalla todos los tweets
func LeerTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Se necesita el id", 400)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Debe enviar la pagina", 400)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Debe enviar una pagina con un numero mayor a 0", 400)
		return
	}
	tweets, status := bd.LeerTweets(ID, pagina)
	if !status {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tweets)
}
