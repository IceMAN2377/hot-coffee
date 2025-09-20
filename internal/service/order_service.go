package service

import (
	"github.com/IceMAN2377/hot-coffee/internal/dal"
	"github.com/IceMAN2377/hot-coffee/internal/models"
)

type OrderService interface {
	CreateOrder(ord *models.CreateOrderMod) (*models.Order, error)
	GetAll() ([]models.Order, error)
	GetOrder(id string) (*models.Order, error)
}

type OrderLogic struct {
	orderRepo dal.OrderRepository
}

func NewOrderLogic(orderRepo dal.OrderRepository) OrderService {
	return &OrderLogic{
		orderRepo: orderRepo,
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
