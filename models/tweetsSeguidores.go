package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//TweetsSeguidores es la estructura para ver los tweets de los seguidores
type TweetsSeguidores struct {
	ID             primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	UsuarioID      string             `bson:"usuario_id" json:"usuario_id,omitempty"`
	UsuarioSeguido string             `bson:"usuarioseguido_id" json:"usuarioseguido_id"`
	Tweet         struct {
		ID      string    `bson:"_id" json:"id,omitempty"`
		Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
		Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
	}
	//UsuarioOID  primitive.ObjectID `bson:"usuario_oid,omitempty" json:"usuario_oid,omitempty"`
	Usuario []*Usuario
}
