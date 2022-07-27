package service_test

import (
	"ass2/entity"
	mock_service "ass2/mock/service"
	"ass2/service"
	"context"
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

		require.Nil(t, res)
	})
	t.Run("success", func(t *testing.T) {
		mockServiceRepo := mock_service.NewMockOrderRepository(ctrl)
		orderService := service.NewOrderSvc(mockServiceRepo)
		order := &entity.Order{
			Customer_name: "test11",
		}
		mockServiceRepo.EXPECT().CreateOrder(context.Background(), order, []entity.Item{}).Return(order)
		res := orderService.CreateOrder(context.Background(), &entity.Order{
			Customer_name: "test11",
		}, []entity.Item{})

		require.Equal(t, order, res)
	})
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Empty id", func(t *testing.T) {
		mockServiceRepo := mock_service.NewMockOrderRepository(ctrl)
		orderService := service.NewOrderSvc(mockServiceRepo)
		order := &entity.Order{
			Customer_name: "Sari",
			Order_id:      0,
		}
		mockServiceRepo.EXPECT().UpdateOrder(context.Background(), order,
			[]entity.Item{
				{
					Item_id:   6,
					Item_code: "Iphone X",
				},
				{
					Item_id:   7,
					Item_code: "Samsung S10+",
				},
			}, 0).Return(nil)

		res := orderService.UpdateOrder(context.Background(), &entity.Order{
			Customer_name: "Sari",
			Order_id:      0,
		}, []entity.Item{
			{
				Item_id:   6,
				Item_code: "Iphone X",
			},
			{
				Item_id:   7,
				Item_code: "Samsung S10+",
			},
		}, 0)

		require.Nil(t, res)
	})

	t.Run("Success", func(t *testing.T) {
		mockServiceRepo := mock_service.NewMockOrderRepository(ctrl)
		orderService := service.NewOrderSvc(mockServiceRepo)
		order := &entity.Order{
			Customer_name: "Sari",
			Order_id:      10,
		}
		mockServiceRepo.EXPECT().UpdateOrder(context.Background(), order,
			[]entity.Item{
				{
					Item_id:   6,
					Item_code: "Iphone X",
				},
				{
					Item_id:   7,
					Item_code: "Samsung S10+",
				},
			}, 10).Return(order)

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

		require.Equal(t, order, res)
	})

}
