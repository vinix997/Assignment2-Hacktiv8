package handler

import (
	"ass2/entity"
	"ass2/service"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type OrderHanlderInterface interface {
	OrderHandler(w http.ResponseWriter, r *http.Request)
}

type OrderHandler struct {
	orderService service.OrderService
	sql          *sql.DB
}

func NewOrderHandler(orderService service.OrderService, sql *sql.DB) OrderHanlderInterface {
	return &OrderHandler{orderService: orderService, sql: sql}
}
func (h *OrderHandler) OrderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["orderId"]
	if r.Method == "POST" {
		h.createOrderHandler(w, r)
	}
	if r.Method == "GET" {
		h.getOrderHandler(w, r)
	}
	if r.Method == "PUT" {
		h.updateOrderHandler(w, r, id)
	}
	if r.Method == "DELETE" {
		h.deleteOrderHandler(w, r, id)
	}

}

func (h *OrderHandler) createOrderHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var order entity.Order
	if err := decoder.Decode(&order); err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	_ = h.orderService.CreateOrder(r.Context(), &entity.Order{
		Customer_name: order.Customer_name,
		Ordered_at:    order.Ordered_at,
	}, order.Items)

	w.Write([]byte("Order created successfully"))
}

func (h *OrderHandler) updateOrderHandler(w http.ResponseWriter, r *http.Request, id interface{}) {
	decoder := json.NewDecoder(r.Body)
	var order entity.Order
	if err := decoder.Decode(&order); err != nil {
		w.Write([]byte("error decoding json body"))
	}

	_ = h.orderService.UpdateOrder(r.Context(), &entity.Order{
		Customer_name: order.Customer_name,
		Ordered_at:    order.Ordered_at,
	}, order.Items, id)

	w.Write([]byte("Order updated successfully"))
}

func (h *OrderHandler) getOrderHandler(w http.ResponseWriter, r *http.Request) {
	query := "SELECT ORDER_ID, CUSTOMER_NAME, ORDERED_AT FROM ORDERS"
	rows, err := h.sql.QueryContext(r.Context(), query)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var orders []*entity.Order
	for rows.Next() {
		var order entity.Order
		if err := rows.Scan(&order.Order_id, &order.Customer_name, &order.Ordered_at); err != nil {
			log.Fatal(err)
		}
		orders = append(orders, &order)
	}

	for i, v := range orders {
		query := "SELECT ITEM_ID, ITEM_CODE, DESCRIPTION, QUANTITY, ORDER_ID FROM ITEMS WHERE ORDER_ID = ?"
		rows, err := h.sql.QueryContext(r.Context(), query, v.Order_id)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var items []entity.Item
		for rows.Next() {
			var item entity.Item
			if err := rows.Scan(&item.Item_id, &item.Item_code, &item.Description, &item.Quantity, &item.Order_id); err != nil {
				log.Fatal(err)
			}
			items = append(items, item)
		}
		orders[i].Items = items
	}

	jsonData, _ := json.Marshal(&orders)
	w.Write(jsonData)
}

func (h *OrderHandler) deleteOrderHandler(w http.ResponseWriter, r *http.Request, id interface{}) {

	query := "DELETE FROM ORDERS WHERE ORDER_ID = ?"
	_, err := h.sql.ExecContext(r.Context(), query, id)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte("Order deleted successfully"))
}
