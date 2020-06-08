package bd

import (
	"context"
	"log"
	"time"

	"github.com/xfchris/gotter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//LeerTweets te muestra la cantidad de tweet
func LeerTweets(ID string, pagina int) ([]*models.Tweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittgo") //selecciono database
	col := db.Collection("tweets")    //selecciono tabla

	var tweets []*models.Tweet

	condicion := bson.M{
		"user_id": ID,
	}

	limit := int64(5)
	opciones := options.Find()
	opciones.SetLimit(limit)
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
	opciones.SetSkip((int64(pagina) - 1)*limit)

	cursor, err := col.Find(ctx, condicion, opciones)

	if err != nil {
		log.Fatal(err.Error())
		return tweets, false
	}

	for cursor.Next(context.TODO()) {

		var reg models.Tweet
		err := cursor.Decode(&reg)
		if err != nil {
			return tweets, true
		}
		tweets = append(tweets, &reg)
	}

	return tweets, true
}
