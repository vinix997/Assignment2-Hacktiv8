package repository

import (
	"ass2/entity"
	"ass2/service"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type repo struct {
	sql *sql.DB
}

func NewOrderRepository(sqlDB *sql.DB) service.OrderRepository {
	return &repo{sql: sqlDB}
}

func (o *repo) CreateOrder(ctx context.Context, order *entity.Order, item []entity.Item) *entity.Order {
	query := "INSERT INTO ORDERS (CUSTOMER_NAME, ORDERED_AT) VALUES(?,?)"
	rows, err := o.sql.ExecContext(ctx, query, order.Customer_name, time.Now())
	if err != nil {
		panic(err)
	}
	insertId, err := rows.LastInsertId()

	itemQuery := "INSERT INTO ITEMS (ITEM_CODE, DESCRIPTION, QUANTITY, ORDER_ID) VALUES (?,?,?,?)"
	for _, v := range item {
		_, err := o.sql.QueryContext(ctx, itemQuery, v.Item_code, v.Description, v.Quantity, insertId)
		if err != nil {
			panic(err)
		}
	}
	return order
}

func (o *repo) UpdateOrder(ctx context.Context, order *entity.Order, item []entity.Item, id interface{}) *entity.Order {
	updateOrderQuery := "UPDATE ORDERS SET CUSTOMER_NAME = ? , ORDERED_AT = ? WHERE ORDER_ID = ? "
	_, err := o.sql.ExecContext(ctx, updateOrderQuery, order.Customer_name, time.Now(), id)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v %v", order, item)
	updateItemQuery := "UPDATE ITEMS SET ITEM_CODE = ?, DESCRIPTION = ?, QUANTITY = ? where ITEM_ID = ? "
	for _, v := range item {
		_, err := o.sql.QueryContext(ctx, updateItemQuery, v.Item_code, v.Description, v.Quantity, v.Item_id)
		if err != nil {
			panic(err)
		}
	}
	return order
}
