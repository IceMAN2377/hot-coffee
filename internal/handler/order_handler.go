package handler

import (
	"encoding/json"
	"github.com/IceMAN2377/hot-coffee/internal/models"
	"github.com/IceMAN2377/hot-coffee/internal/service"
	"log"
	"net/http"
)

type OrderHandler struct {
	orderServ service.OrderService
}

func NewOrderHandler(orderServ service.OrderService) *OrderHandler {
	return &OrderHandler{
		orderServ: orderServ,
	}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var ord models.CreateOrderMod

	if err := json.NewDecoder(r.Body).Decode(&ord); err != nil {
		log.Println("failed to decode order")
	}

	order, err := h.orderServ.CreateOrder(&ord)
	if err != nil {
		log.Println("failed to send order to service layer")
	}
	json.NewEncoder(w).Encode(order)

}

func (h *OrderHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	orders, err := h.orderServ.GetAll()
	if err != nil {
		log.Printf("failed to get all orders: %v", err)
	}
	json.NewEncoder(w).Encode(orders)
}
