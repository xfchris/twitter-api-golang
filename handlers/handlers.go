package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/xfchris/gotter/middle"
	"github.com/xfchris/gotter/routers"
)

//Manejadores son las rutas
func Manejadores() {
	router := mux.NewRouter()
	router.HandleFunc("/", routers.Index).Methods("GET")

	router.HandleFunc("/registro", middle.ChequeoDB(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middle.ChequeoDB(routers.Login)).Methods("POST")
	router.HandleFunc("/ver-perfil", middle.ChequeoDB(middle.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/editar-perfil", middle.ChequeoDB(middle.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/crear-tweet", middle.ChequeoDB(middle.ValidoJWT(routers.CrearTweet))).Methods("POST")
	router.HandleFunc("/leer-tweet", middle.ChequeoDB(middle.ValidoJWT(routers.LeerTweets))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
