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
	router.HandleFunc("/registro", middle.ChequeoDB(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middle.ChequeoDB(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middle.ChequeoDB(middle.ChequeoDB(routers.VerPerfil))).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
