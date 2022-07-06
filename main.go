package main

import (
	"ass2/handler"
	"ass2/repository"
	"ass2/service"
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const PORT = ":8080"

func main() {
	orderRepo := repository.NewOrderRepository(Connection())
	r := mux.NewRouter()

	orderService := service.NewOrderSvc(orderRepo)

	createOrderHandler := handler.NewOrderHandler(orderService, Connection())

	r.HandleFunc("/orders", createOrderHandler.OrderHandler)
	r.HandleFunc("/orders/{orderId}", createOrderHandler.OrderHandler)
	http.Handle("/", r)
	http.ListenAndServe(PORT, nil)
}

func Connection() *sql.DB {
	database, err := sql.Open("mysql", "root:admin123@tcp(localhost:3306)/orders_by?parseTime=true")
	if err != nil {
		panic(err)
	}
	return database
}
