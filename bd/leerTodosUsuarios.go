package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/xfchris/gotter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//LeerUsuarios Muestra lista de usuarios dependiendo de atributos
func LeerUsuarios(ID string, page int, search string, tipo string) ([]*models.Usuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittgo") //selecciono database
	col := db.Collection("usuarios")

	var results []*models.Usuario
	limit := int64(5)

	findOptions := options.Find()
	findOptions.SetSkip((int64(page) - 1) * limit)
	findOptions.SetLimit(limit)
	findOptions.SetProjection(
		bson.M{
			"email" : 0,
			"password" : 0,
			"banner" : 0,
			"biografia" : 0,
			"ubicacion" : 0,
			"sitioWeb" : 0,
		},
	)

	query := bson.M{
		"nombres": bson.M{
			"$regex": `(?i)` + search,
		},
	}

	cur, err := col.Find(ctx, query, findOptions)
	
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	var encontrado, incluir bool
	for cur.Next(ctx) {
		var s models.Usuario
		err = cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		var r models.UsuarioSeguido
		r.UsuarioID = ID
		r.UsuarioSeguido = s.ID.Hex()

		incluir = false

		encontrado, err = ConsultoSeguidor(r)
		if tipo == "new" && encontrado == false {
			incluir = true
		} else if tipo == "follow" && encontrado == true {
			incluir = true
		}

		if r.UsuarioSeguido == r.UsuarioID {
			incluir = false
		}

		if incluir {
			// s.Password = ""
			// s.Biografia = ""
			// s.SitioWeb = ""
			// s.Ubicacion = ""
			// s.Banner = ""
			// s.Email = ""

			results = append(results, &s)
		}
	}

	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	cur.Close(ctx)

	return results, true
}
