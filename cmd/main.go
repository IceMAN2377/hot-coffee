package main

import (
	"fmt"
	"github.com/IceMAN2377/hot-coffee/internal/dal"
	"github.com/IceMAN2377/hot-coffee/internal/handler"
	"github.com/IceMAN2377/hot-coffee/internal/service"
	"log"
	"net/http"
	"os"
)

func main() {

	filePath := "orders.json"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("failed to create orders json")
	}
	defer file.Close()

	mux := http.NewServeMux()

	orderRepo := dal.NewOrderStore(filePath)
	orderServ := service.NewOrderLogic(orderRepo)
	orderHand := handler.NewOrderHandler(orderServ)

	mux.HandleFunc("POST /orders", orderHand.CreateOrder)

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("error connecting")
	}

}
