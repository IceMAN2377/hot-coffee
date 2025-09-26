package handler

import (
	"github.com/IceMAN2377/hot-coffee/internal/service"
	"net/http"
)

func RegisterOrder(router *http.ServeMux, service service.OrderService) {
	handler := NewOrderHandler(service)

	router.HandleFunc("POST /orders", handler.CreateOrder)
	router.HandleFunc("GET /orders", handler.GetAll)
	router.HandleFunc("GET /orders/{id}", handler.GetOrder)
	router.HandleFunc("PUT /orders/{id}", handler.UpdateOrder)
	router.HandleFunc("DELETE /orders/{id}", handler.DeleteOrder)
	router.HandleFunc("POST /orders/{id}/close", handler.CloseOrder)
}
