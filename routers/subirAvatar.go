package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/xfchris/gotter/bd"
	"github.com/xfchris/gotter/middle"
	"github.com/xfchris/gotter/models"
)

//SubirAvatar sube un avatar al sistema
func SubirAvatar(w http.ResponseWriter, r *http.Request) {
	subirAvatarBanner("avatar", w, r)
}

//SubirBanner sube un banner al sistema
func SubirBanner(w http.ResponseWriter, r *http.Request) {
	subirAvatarBanner("banner", w, r)
}

func subirAvatarBanner(img string, w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile(img)
	if err != nil {
		http.Error(w, "Error al capturar imagen del formulario: "+err.Error(), 400)
		return
	}
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/" + img + "s/" + middle.IDUsuario + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir imagen: "+err.Error(), 400)
		return
	}
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar imagen: "+err.Error(), 400)
		return
	}

	var usuario models.Usuario
	var status bool
	if img == "avatar" {
		usuario.Avatar = middle.IDUsuario + "." + extension
	} else {
		usuario.Banner = middle.IDUsuario + "." + extension
	}
	status, err = bd.ModificarRegistro(usuario, middle.IDUsuario)

	if err != nil || !status {
		http.Error(w, "Error al grabar "+img+"s en la DB "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
