package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/xfchris/gotter/bd"
	"github.com/xfchris/gotter/middle"
	"github.com/xfchris/gotter/models"
)

func CrearTweet(w http.ResponseWriter, r *http.Request) {
	var tweet models.Tweet
	err := json.NewDecoder(r.Body).Decode(&tweet)

	tweet.UserID = middle.IDUsuario
	tweet.Fecha = time.Now()

	_, status, err := bd.InsertarTweet(tweet)
	if err != nil {
		http.Error(w, "Error al crear tweet: "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se insertó el tweet", 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tweet)
}
