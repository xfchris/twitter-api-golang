package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Tweet es el modelo de un tweeter
type Tweet struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID  string             `bson:"user_id" json:"user_id,omitempty"`
	Mensaje string             `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time          `bson:"fecha" json:"fecha,omitempty"`
}
