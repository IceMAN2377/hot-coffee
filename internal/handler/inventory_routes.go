package handler

import (
	"github.com/IceMAN2377/hot-coffee/internal/service"
	"net/http"
)

func RegisterInventory(router *http.ServeMux, invServ service.InventoryService) {
	handler := NewInventoryHandler(invServ)

	router.HandleFunc("POST /inventory", handler.AddItem)
	router.HandleFunc("GET /inventory", handler.GetItems)
	router.HandleFunc("GET /inventory/{id}", handler.GetItem)
	router.HandleFunc("PUT /inventory/{id}", handler.UpdateItem)
	router.HandleFunc("DELETE /inventory/{id}", handler.DeleteItem)
}
