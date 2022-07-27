package service_test

import (
	"ass2/entity"
	mock_service "ass2/mock/service"
	"ass2/service"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestNewOrderSvc(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Run("Test initiating new order service", func(t *testing.T) {
		mockOrderRepo := mock_service.NewMockOrderRepository(ctrl)
		orderService := service.NewOrderSvc(mockOrderRepo)
		require.NotNil(t, orderService)
	})
}

func TestOrderSvc_CreateOrder(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Empty order", func(t *testing.T) {
		mockServiceRepo := mock_service.NewMockOrderRepository(ctrl)
		orderService := service.NewOrderSvc(mockServiceRepo)
		res := orderService.CreateOrder(context.Background(), &entity.Order{
			Customer_name: "",
		}, []entity.Item{})
		if res == nil {
			err := errors.New("Customer name cannot be empty")
			require.Error(t, err)
			require.Equal(t, errors.New("Customer name cannot be empty"), err)
		}
		require.Nil(t, res)
	})
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Empty id", func(t *testing.T) {
		mockServiceRepo := mock_service.NewMockOrderRepository(ctrl)
		orderService := service.NewOrderSvc(mockServiceRepo)
		res := orderService.UpdateOrder(context.Background(), &entity.Order{
			Customer_name: "Sari",
			Order_id:      10,
		}, []entity.Item{
			{
				Item_id:   6,
				Item_code: "Iphone X",
			},
			{
				Item_id:   7,
				Item_code: "Samsung S10+",
			},
		}, 10)

		require.Nil(t, res)
	})
}
