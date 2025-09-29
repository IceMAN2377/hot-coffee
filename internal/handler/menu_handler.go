package handler

import (
	"encoding/json"
	"github.com/IceMAN2377/hot-coffee/internal/models"
	"github.com/IceMAN2377/hot-coffee/internal/service"
	"net/http"
)

type MenuHandler struct {
	menuServ service.MenuService
}

func NewMenuHandler(menuServ service.MenuService) *MenuHandler {
	return &MenuHandler{
		menuServ: menuServ,
	}
}

func (h *MenuHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	var item *models.MenuItem

	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	menuPosition, _ := h.menuServ.AddItem(item)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(menuPosition); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h *MenuHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	items, _ := h.menuServ.GetItems()

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(items); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

}

func (h *MenuHandler) GetItem(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	item, err := h.menuServ.GetItem(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h *MenuHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var item *models.MenuItem

	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	updItem, err := h.menuServ.UpdateItem(id, item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(updItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h *MenuHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	delItem := h.menuServ.DeleteItem(id)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(delItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
