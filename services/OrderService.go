package services

import "delivery/repositories"

type OrderService struct {
	repository *repositories.OrderRepository
}

func NewOrderService(repository *repositories.OrderRepository) *OrderService {
	return &OrderService{
		repository: repository,
	}
}
