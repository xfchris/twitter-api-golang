package main

import (
	"log"

	"github.com/xfchris/gotter/bd"
	//"github.com/xfchris/gotter/handlers"
)

func main() {
	if bd.CheckearConexion() == 0 {
		log.Println("No se pudo conectar")
	}
	//handlers.Manejadores()
}
