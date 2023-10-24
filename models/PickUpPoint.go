package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PickUpPoint struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	LocationX float64            `bson:"locationX"`
	LocationY float64            `bson:"locationY"`
}
