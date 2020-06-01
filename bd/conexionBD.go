package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN es el objeto de conexion
var MongoCN = ConectarDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://mongochris:mongochris@cluster0-yzrbk.mongodb.net/test?retryWrites=true&w=majority")

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
