package bd

import (
	"context"
	"time"

	"github.com/xfchris/gotter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertarTweet inserta un tweet en el registro
func InsertarTweet(t models.Tweet) (primitive.ObjectID, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittgo") //selecciono database
	col := db.Collection("tweets")    //selecciono tabla

	registro := bson.M{
		"user_id": t.UserID,
		"mensaje": t.Mensaje,
		"fecha":   t.Fecha,
	}

	result, err := col.InsertOne(ctx, registro)
	if err != nil {
		return primitive.NilObjectID, false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID, true, nil
}
