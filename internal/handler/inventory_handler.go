package handler

import (
	"encoding/json"
	"github.com/IceMAN2377/hot-coffee/internal/models"
	"github.com/IceMAN2377/hot-coffee/internal/service"
	"net/http"
)

type InventoryHandler struct {
	invServ service.InventoryService
}

func NewInventoryHandler(invServ service.InventoryService) *InventoryHandler {
	return &InventoryHandler{
		invServ: invServ,
	}
}

func (h *InventoryHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	var item *models.InventoryItem

	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	menuPosition, _ := h.invServ.AddItem(item)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(menuPosition); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *InventoryHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	items, _ := h.invServ.GetItems()

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(items); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func (h *InventoryHandler) GetItem(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	item, err := h.invServ.GetItem(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *InventoryHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var item *models.InventoryItem

	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	updItem, err := h.invServ.UpdateItem(id, item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(updItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *InventoryHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	err := h.invServ.DeleteItem(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
