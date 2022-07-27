package main

import (
	"ass2/handler"
	"ass2/repository"
	"ass2/service"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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

func getEnvVar(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func Connection() *sql.DB {
	user := getEnvVar("DB_USERNAME")
	password := getEnvVar("PASSWORD")
	db := getEnvVar("DB_NAME")
	database, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?parseTime=true", user, password, db))
	if err != nil {
		panic(err)
	}
	return database
}
