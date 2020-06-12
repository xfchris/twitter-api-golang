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
	router.HandleFunc("/borrar-tweet", middle.ChequeoDB(middle.ValidoJWT(routers.BorrarTweet))).Methods("DELETE")

	router.HandleFunc("/subir-avatar", middle.ChequeoDB(middle.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/subir-banner", middle.ChequeoDB(middle.ValidoJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtener-avatar", middle.ChequeoDB(routers.ObtenerAvatar)).Methods("GET")
	router.HandleFunc("/obtener-banner", middle.ChequeoDB(routers.ObtenerBanner)).Methods("GET")

	router.HandleFunc("/seguir-usuario", middle.ChequeoDB(middle.ValidoJWT(routers.SeguirUsuario))).Methods("POST")
	router.HandleFunc("/borrar-seguidor", middle.ChequeoDB(middle.ValidoJWT(routers.BorrarSeguidor))).Methods("DELETE")
	router.HandleFunc("/consultar-seguidor", middle.ChequeoDB(middle.ValidoJWT(routers.ConsultarSeguidor))).Methods("GET")

	router.HandleFunc("/leer-usuarios", middle.ChequeoDB(middle.ValidoJWT(routers.LeerUsuarios))).Methods("GET")
	router.HandleFunc("/leer-seguidores", middle.ChequeoDB(middle.ValidoJWT(routers.LeerTweetsSeguidores))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
