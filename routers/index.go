package routers

import (
	"fmt"
	"net/http"
)

//Index principal
func Index(w http.ResponseWriter, r *http.Request) {

	//w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "<h1>Hello go</h1>")
}
