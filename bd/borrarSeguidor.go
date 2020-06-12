package bd

import (
	"context"
	"time"

	"github.com/xfchris/gotter/models"
)

//BorrarSeguidor borra una relacion entre usuarios
func BorrarSeguidor(t models.UsuarioSeguido) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittgo") //selecciono database
	col := db.Collection("usuario_seguido")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
