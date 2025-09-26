package service

import (
	"github.com/IceMAN2377/hot-coffee/internal/dal"
	"github.com/IceMAN2377/hot-coffee/internal/models"
)

type MenuService interface {
	AddItem(item *models.MenuItem) (*models.MenuItem, error)
	GetItems() ([]models.MenuItem, error)
	GetItem(id string) (*models.MenuItem, error)
	UpdateItem(id string, item *models.MenuItem) (*models.MenuItem, error)
}

type MenuLogic struct {
	menuRepo dal.MenuRepository
}

func NewMenuLogic(menuRepo dal.MenuRepository) MenuService {
	return &MenuLogic{
		menuRepo: menuRepo,
	}
}

func (s *MenuLogic) AddItem(item *models.MenuItem) (*models.MenuItem, error) {
	return s.menuRepo.AddItem(item)
}

func (s *MenuLogic) GetItems() ([]models.MenuItem, error) {
	return s.menuRepo.GetItems()
}

func (s *MenuLogic) GetItem(id string) (*models.MenuItem, error) {
	return s.menuRepo.GetItem(id)
}

func (s *MenuLogic) UpdateItem(id string, item *models.MenuItem) (*models.MenuItem, error) {
	return s.menuRepo.UpdateItem(id, item)
}
