package controllers

import (
	"delivery/services"
)

type OrderController struct {
	service *services.OrderService
}

func NewOrderController(service *services.OrderService) *OrderController {
	return &OrderController{
		service: service,
	}
}

func (c *OrderController) AddOrderToDelivery(order *Order) *Response {
	if order.Item == "" || order.UserName == "" || order.UserPhone == "" || order.Address == "" || order.UserLocationX == 0 || order.UserLocationY == 0 || order.PickUpPointX == 0 || order.PickUpPointY == 0 {
		return &Response{
			Status:  "failed",
			Message: "invalid order data",
		}
	}
}
