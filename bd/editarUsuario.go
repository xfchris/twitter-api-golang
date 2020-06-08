package bd

import (
	"context"
	"time"

	"github.com/xfchris/gotter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ModificarRegistro modifica un usuario en base de datos, devuele un booleano si fue exitoso o no
func ModificarRegistro(u models.Usuario, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittgo") //selecciono database
	col := db.Collection("usuarios")  //selecciono tabla

	updString := bson.M{
		"$set": generarOBJactualizado(u),
	}
	objID, _ := primitive.ObjectIDFromHex(ID)

	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filtro, updString)

	// if err != nil {
	// 	return false, err
	// }
	// return true, nil
	return !(err != nil), err
}

func generarOBJactualizado(u models.Usuario) map[string]interface{} {
	registro := make(map[string]interface{})

	if len(u.Nombres) > 0 {
		registro["nombres"] = u.Nombres
	}
	if len(u.Apellidos) > 0 {
		registro["apellidos"] = u.Apellidos
	}
	registro["fechaNacimiento"] = u.FechaNacimiento

	if len(u.Avatar) > 0 {
		registro["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		registro["banner"] = u.Banner
	}
	if len(u.Biografia) > 0 {
		registro["biografia"] = u.Biografia
	}
	if len(u.Ubicacion) > 0 {
		registro["ubicacion"] = u.Ubicacion
	}
	if len(u.SitioWeb) > 0 {
		registro["sitioWeb"] = u.SitioWeb
	}
	return registro
}
