package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type BItem struct {
	ID            primitive.ObjectID `bson:"_id"`
	ItemID        primitive.ObjectID `bson:"item_id"`
	UserID        primitive.ObjectID `bson:"user_id"`
	PickUpPointID primitive.ObjectID `bson:"pick_up_point_id"`

	MetaInfo    string    `bson:"meta_info"`
	Status      string    `bson:"status"`
	DeliveredAt time.Time `bson:"delivered_at"`
}

type Order struct {
	ID          primitive.ObjectID
	Item        Item
	User        User
	PickUpPoint PickUpPoint

	MetaInfo    string
	Status      string
	DeliveredAt time.Time
}
