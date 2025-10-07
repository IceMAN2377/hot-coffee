package service

import (
	"github.com/IceMAN2377/hot-coffee/internal/dal"
	"github.com/IceMAN2377/hot-coffee/internal/models"
)

type InventoryService interface {
	AddItem(item *models.InventoryItem) (*models.InventoryItem, error)
	GetItems() ([]models.InventoryItem, error)
	GetItem(id string) (*models.InventoryItem, error)
	UpdateItem(id string, item *models.InventoryItem) (*models.InventoryItem, error)
	DeleteItem(id string) error
}

type InventoryLogic struct {
	invRepo dal.InventoryRepository
}

func NewInventoryLogic(invRepo dal.InventoryRepository) *InventoryLogic {
	return &InventoryLogic{
		invRepo: invRepo,
	}
}

func (s *InventoryLogic) AddItem(item *models.InventoryItem) (*models.InventoryItem, error) {
	return s.invRepo.AddItem(item)
}

func (s *InventoryLogic) GetItems() ([]models.InventoryItem, error) {
	return s.invRepo.GetItems()
}

func (s *InventoryLogic) GetItem(id string) (*models.InventoryItem, error) {
	return s.invRepo.GetItem(id)
}

func (s *InventoryLogic) UpdateItem(id string, item *models.InventoryItem) (*models.InventoryItem, error) {
	return s.invRepo.UpdateItem(id, item)
}

func (s *InventoryLogic) DeleteItem(id string) error {
	return s.invRepo.DeleteItem(id)
}
