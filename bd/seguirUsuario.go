package bd

import (
	"context"
	"time"

	"github.com/xfchris/gotter/models"
)

//SeguirUsuario crea relacion entre usuarios
func SeguirUsuario(t models.UsuarioSeguido) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittgo") //selecciono database
	col := db.Collection("usuario_seguido")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
