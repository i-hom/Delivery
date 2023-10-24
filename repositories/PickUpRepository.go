package repositories

import (
	"delivery/models"
	"delivery/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PickUpPointRepository struct {
	db             *storage.Storage
	collectionName string
}

func NewPickUpPointRepository(db *storage.Storage) *PickUpPointRepository {
	return &PickUpPointRepository{
		db:             db,
		collectionName: "PickUpPoints",
	}
}

func (r *PickUpPointRepository) GetPickUpPoint(ID primitive.ObjectID) (models.PickUpPoint, error) {
	var pickUpPoint models.PickUpPoint
	err := r.db.Find(r.collectionName, bson.M{"_id": ID}, &pickUpPoint)
	return pickUpPoint, err
}
