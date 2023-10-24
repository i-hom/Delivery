package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Phone     string             `bson:"phone"`
	Address   string             `bson:"address"`
	Role      string             `bson:"role"`
	LocationX float64            `bson:"locationX"`
	LocationY float64            `bson:"locationY"`
}
