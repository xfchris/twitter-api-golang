package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/xfchris/gotter/models"
	"go.mongodb.org/mongo-driver/bson"
)

//LeerTweetsSeguidores te devuelve los tweets de mis seguidores
func LeerTweetsSeguidores(ID string, page int) ([]*models.TweetsSeguidores, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittgo")
	col := db.Collection("usuario_seguido")
	limit := 5
	skip := (page - 1) * limit

	condiciones := make([]bson.M, 0)

	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuario_id": ID}})
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "tweets",
			"localField":   "usuarioseguido_id",
			"foreignField": "user_id",
			"as":           "tweet",
		},
	})

	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})

	//Mucho a muchos con usuarios
	condiciones = append(condiciones, bson.M{
		"$addFields": bson.M{"usuario_oid": bson.M{"$toObjectId": "$usuarioseguido_id"}},
	})
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "usuarios",
			"localField":   "usuario_oid",
			"foreignField": "_id",
			"as":           "usuario",
		},
	})
	//Fin mucho a muchos con usuarios
	//Oculto algunos campos
	condiciones = append(condiciones, bson.M{
		"$project": bson.M{
			"usuario_id":              0,
			"usuario.password":        0,
			"usuario.sitioWeb":        0,
			"usuario.biografia":       0,
			"usuario._id":             0,
			"tweets.user_id":          0,
		},
	})

	condiciones = append(condiciones, bson.M{"$sort": bson.M{"tweets.fecha": -1}})
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": limit})

	cur, err := col.Aggregate(ctx, condiciones)

	var results []*models.TweetsSeguidores
	err = cur.All(ctx, &results)
	if err != nil {
		fmt.Println("Error en cursor: " + err.Error())
		return results, false
	}
	return results, true
}
