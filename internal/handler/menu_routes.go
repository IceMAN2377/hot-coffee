package handler

import (
	"github.com/IceMAN2377/hot-coffee/internal/service"
	"net/http"
)

func RegisterMenu(router *http.ServeMux, menuServ service.MenuService) {
	handler := NewMenuHandler(menuServ)

	router.HandleFunc("POST /menu", handler.AddItem)
	router.HandleFunc("GET /menu", handler.GetItems)
	router.HandleFunc("GET /menu/{id}", handler.GetItem)
	router.HandleFunc("PUT /menu/{id}", handler.UpdateItem)
}
