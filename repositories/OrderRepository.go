package repositories

import (
	"delivery/models"
	"delivery/storage"
	"go.mongodb.org/mongo-driver/bson"
)

type OrderRepository struct {
	db             *storage.Storage
	collectionName string
}

func NewOrderRepository(db *storage.Storage) *OrderRepository {
	return &OrderRepository{
		db:             db,
		collectionName: "Orders",
	}
}

func (r *OrderRepository) Get() ([]models.Order, error) {
	var orders []models.Order
	err := r.db.FindAll(r.collectionName, bson.M{}, &orders)
	return orders, err
}

func (r *OrderRepository) Add(Order models.Order) error {
	return r.db.Insert(r.collectionName, Order)
}
