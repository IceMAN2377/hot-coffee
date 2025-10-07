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
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	order, err := h.orderServ.CreateOrder(&ord)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(order); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (h *OrderHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	orders, err := h.orderServ.GetAll()
	if err != nil {
		log.Printf("failed to get all orders: %v", err)
	}
	if err := json.NewEncoder(w).Encode(orders); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	order, err := h.orderServ.GetOrder(id)
	if err != nil {
		log.Printf("failed to get order by id: %v", err)
	}
	if err := json.NewEncoder(w).Encode(order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h *OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var items []models.OrderItem

	if err := json.NewDecoder(r.Body).Decode(&items); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	updOrder, err := h.orderServ.UpdateOrder(id, items)
	if err != nil {
		log.Printf("failed to update order by id: %v", err)
	}
	if err := json.NewEncoder(w).Encode(updOrder); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h *OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	delOrder := h.orderServ.DeleteOrder(id)

	if err := json.NewEncoder(w).Encode(delOrder); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h *OrderHandler) CloseOrder(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	closeOrder := h.orderServ.CloseOrder(id)

	if err := json.NewEncoder(w).Encode(closeOrder); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
