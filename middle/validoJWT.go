package middle

import (
	"net/http"
)

func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := ProcesarToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en token "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
