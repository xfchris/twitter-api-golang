package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/xfchris/gotter/bd"
)

//ObtenerAvatar obtiene el avatar de un usuario cualquiera
func ObtenerAvatar(w http.ResponseWriter, r *http.Request) {
	obtenerAvatarBanner("avatar", w, r)
}

//ObtenerBanner obtiene el banner de un usuario cualquiera
func ObtenerBanner(w http.ResponseWriter, r *http.Request) {
	obtenerAvatarBanner("banner", w, r)
}

func obtenerAvatarBanner(img string, w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Es necesario el id del usuario: ", 400)
		return
	}

	usuario, err := bd.BuscarPerfil(ID)
	if err != nil {
		http.Error(w, "Usuario no encontrado: "+err.Error(), 400)
		return
	}
	var filename string
	if img == "avatar" {
		filename = "uploads/" + img + "s/" + usuario.Avatar
	} else {
		filename = "uploads/" + img + "s/" + usuario.Banner
	}
	file, err := os.Open(filename)
	if err != nil {
		http.Error(w, "Imagen no encontrada: "+err.Error(), 400)
		return
	}
	//pongo el archivo en la respuesta
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Error al copiar imagen: "+err.Error(), 400)
		return
	}
}
