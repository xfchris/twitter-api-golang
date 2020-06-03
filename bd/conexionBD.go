package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN es el objeto de conexion
var MongoCN = ConectarDB()

//var con = "mongodb://127.0.0.1:27017"
var con = "mongodb://mongochris:mongochris@cluster0-shard-00-00-yzrbk.mongodb.net:27017,cluster0-shard-00-01-yzrbk.mongodb.net:27017,cluster0-shard-00-02-yzrbk.mongodb.net:27017/test?ssl=true&replicaSet=Cluster0-shard-0&authSource=admin&retryWrites=true&w=majority"
var clientOptions = options.Client().ApplyURI(con)

/*ConectarDB permite Conectar a la base de datos */
func ConectarDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion exitosa")
	return client
}

//CheckearConexion checkea si la conexion es exitosa o no mediante un ping
func CheckearConexion() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
