package service

import (
	"github.com/IceMAN2377/hot-coffee/internal/dal"
	"github.com/IceMAN2377/hot-coffee/internal/models"
)

type OrderService interface {
	CreateOrder(ord *models.CreateOrderMod) (*models.Order, error)
	GetAll() ([]models.Order, error)
	GetOrder(id string) (*models.Order, error)
	UpdateOrder(id string, items []models.OrderItem) (*models.Order, error)
	DeleteOrder(id string) error
	CloseOrder(id string) error
}

type OrderLogic struct {
	orderRepo dal.OrderRepository
	menuServ  MenuService
	invServ   InventoryService
}

func NewOrderLogic(orderRepo dal.OrderRepository, menuServ MenuService, invServ InventoryService) OrderService {
	return &OrderLogic{
		orderRepo: orderRepo,
		menuServ:  menuServ,
		invServ:   invServ,
	}
}

func (s *OrderLogic) CreateOrder(ord *models.CreateOrderMod) (*models.Order, error) {
	return s.orderRepo.CreateOrder(ord)
}

func (s *OrderLogic) GetAll() ([]models.Order, error) {
	return s.orderRepo.GetAll()
}

func (s *OrderLogic) GetOrder(id string) (*models.Order, error) {
	return s.orderRepo.GetOrder(id)
}

func (s *OrderLogic) UpdateOrder(id string, items []models.OrderItem) (*models.Order, error) {
	return s.orderRepo.UpdateOrder(id, items)
}

func (s *OrderLogic) DeleteOrder(id string) error {
	return s.orderRepo.DeleteOrder(id)
}

func (s *OrderLogic) CloseOrder(id string) error {
	return s.orderRepo.CloseOrder(id)
}
