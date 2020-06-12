package bd

import (
	"context"
	"time"

	"github.com/xfchris/gotter/models"
	"go.mongodb.org/mongo-driver/bson"
)

//ConsultoSeguidor Me confirma si estoy relacionado con un suario
func ConsultoSeguidor(t models.UsuarioSeguido) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittgo") //selecciono database
	col := db.Collection("usuario_seguido")

	condicion := bson.M{
		"usuario_id":        t.UsuarioID,
		"usuarioseguido_id": t.UsuarioSeguido,
	}

	var res models.UsuarioSeguido
	err := col.FindOne(ctx, condicion).Decode(&res)

	if err != nil {
		return false, err
	}
	return true, nil
}
