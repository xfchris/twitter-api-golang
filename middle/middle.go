package middle

import (
	"net/http"

	"github.com/xfchris/gotter/bd"
)

//ChequeoDB comprueba si la base de datos esta conectada
func ChequeoDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckearConexion() == 0 {
			http.Error(w, "Conexion perdida con DB", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
