package repositories

import (
	"delivery/models"
	"delivery/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository struct {
	db             *storage.Storage
	collectionName string
}

func NewUserRepository(db *storage.Storage) *UserRepository {
	return &UserRepository{
		db:             db,
		collectionName: "Users",
	}
}

func (r *UserRepository) GetUser(ID primitive.ObjectID) (models.User, error) {
	var user models.User
	err := r.db.Find(r.collectionName, bson.M{"_id": ID}, &user)
	return user, err
}
