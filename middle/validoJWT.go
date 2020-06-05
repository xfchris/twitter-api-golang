package middle

import (
	"net/http"

	"github.com/xfchris/gotter/routers"
)

func validoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesarToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en token "+err.Error(), http.StatusBadRequest)
		}
		next.ServeHTTP(w, r)
	}
}
