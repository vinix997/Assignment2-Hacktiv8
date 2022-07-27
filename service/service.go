package service

import (
	"ass2/entity"
	"context"
	"errors"
)

type OrderService interface {
	CreateOrder(ctx context.Context, order *entity.Order, items []entity.Item) *entity.Order
	UpdateOrder(ctx context.Context, order *entity.Order, items []entity.Item, id interface{}) *entity.Order
}

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *entity.Order, items []entity.Item) *entity.Order
	UpdateOrder(ctx context.Context, order *entity.Order, items []entity.Item, id interface{}) *entity.Order
}

type OrderSvc struct {
	orderRepo OrderRepository
}

func NewOrderSvc(orderRepo OrderRepository) OrderService {
	return &OrderSvc{
		orderRepo: orderRepo,
	}
}

func (u *OrderSvc) CreateOrder(ctx context.Context, order *entity.Order, items []entity.Item) *entity.Order {
	if err := validateOrder(order, items); err != nil {
		return nil
	}
	return u.orderRepo.CreateOrder(ctx, order, items)
}
func (u *OrderSvc) UpdateOrder(ctx context.Context, order *entity.Order, items []entity.Item, id interface{}) *entity.Order {
	return u.orderRepo.UpdateOrder(ctx, order, items, id)
}

func validateOrder(order *entity.Order, items []entity.Item) error {
	if order == nil {
		return errors.New("Order cannot be empty")
	}
	if order.Customer_name == "" {
		return errors.New("Customer name cannot be empty")
	}
	if items == nil {
		return errors.New("Item cannot be empty")
	}
	return nil
}
