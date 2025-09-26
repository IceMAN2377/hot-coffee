package main

import (
	"fmt"
	"github.com/IceMAN2377/hot-coffee/internal/dal"
	"github.com/IceMAN2377/hot-coffee/internal/handler"
	"github.com/IceMAN2377/hot-coffee/internal/service"
	"net/http"
)

func main() {

	//logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	filePathOrder := "orders.json"
	//file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	//if err != nil {
	//	log.Fatal("failed to create orders json")
	//}
	//defer file.Close()

	mux := http.NewServeMux()

	orderRepo := dal.NewOrderStore(filePathOrder)
	orderServ := service.NewOrderLogic(orderRepo)
	handler.RegisterOrder(mux, orderServ)

	filePathMenu := "menu.json"

	menuRepo := dal.NewMenuStore(filePathMenu)
	menuServ := service.NewMenuLogic(menuRepo)
	handler.RegisterMenu(mux, menuServ)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("error connecting")
	}

}
