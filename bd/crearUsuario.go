package bd

import (
	"context"
	"time"

	"github.com/xfchris/gotter/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//CrearUsuario inserta un nuevo usuario ne la Base de datos
func CrearUsuario(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittgo")             //selecciono database
	col := db.Collection("usuarios")              //selecciono tabla
	u.Password, _ = EncriptarPassword(u.Password) //encripto password

	result, err := col.InsertOne(ctx, u) //inserto registro en base de dato
	if err != nil {
		return "", false, err
	}
	objID := result.InsertedID.(primitive.ObjectID) //obtengo id

	return objID.String(), true, nil //retorno
}
