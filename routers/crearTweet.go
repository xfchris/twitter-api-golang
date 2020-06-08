package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/xfchris/gotter/bd"
	"github.com/xfchris/gotter/middle"
	"github.com/xfchris/gotter/models"
)

//CrearTweet crea un tweet basado en campos enviados por post
func CrearTweet(w http.ResponseWriter, r *http.Request) {
	var tweet models.Tweet
	err := json.NewDecoder(r.Body).Decode(&tweet)

	tweet.UserID = middle.IDUsuario
	tweet.Fecha = time.Now()

	id, status, err := bd.InsertarTweet(tweet)
	if err != nil {
		http.Error(w, "Error al crear tweet: "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se insertó el tweet", 400)
		return
	}
	// id = strings.Replace(id, "ObjectID(\"", "", 1)
	// id = strings.Replace(id, "\")", "", 1)
	// objID, _ := primitive.ObjectIDFromHex(id)
	// tweet.ID = objID
	tweet.ID = id

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tweet)
}
