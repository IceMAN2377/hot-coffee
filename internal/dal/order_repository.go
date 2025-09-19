package dal

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/IceMAN2377/hot-coffee/internal/models"
	"log"
	"os"
	"sync"
	"time"
)

type OrderRepository interface {
	CreateOrder(ord *models.CreateOrderMod) (models.Order, error)
	GetAll() ([]models.Order, error)
	GetOrder(id string) (models.Order, error)
	UpdateOrder(id string) (models.Order, error)
	DeleteOrder(id string) error
	CloseOrder(id string) error
}

type OrderStore struct {
	//logger   slog.Logger
	filePath string
	mutex    sync.RWMutex
	orders   []models.Order
	nextID   int
}

func NewOrderStore(filePath string) *OrderStore {
	o := &OrderStore{
		//logger:   *logger,
		filePath: filePath,
		nextID:   1,
	}
	err := o.LoadFromFile()
	if err != nil {
		log.Printf("error loading: %v", err)
	}
	return o
}

func (o *OrderStore) CreateOrder(ord *models.CreateOrderMod) (models.Order, error) {
	id := fmt.Sprintf("%d", o.nextID)
	status := "open"
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	order := models.Order{
		ID:           id,
		CustomerName: ord.CustomerName,
		Items:        ord.Items,
		Status:       status,
		CreatedAt:    currentTime,
	}

	o.orders = append(o.orders, order)

	if err := o.SaveToFile(o.orders); err != nil {
		log.Fatal("failed to save to order json")
	}
	o.nextID++

	return order, nil
}

func (o *OrderStore) GetAll() ([]models.Order, error) {

	err := o.LoadFromFile()
	if err != nil {
		log.Printf("error loading: %v", err)
	}

	return o.orders, nil
}

func (o *OrderStore) GetOrder(id string) (models.Order, error) {
	return models.Order{}, nil
}

func (o *OrderStore) UpdateOrder(id string) (models.Order, error) {
	return models.Order{}, nil
}

func (o *OrderStore) DeleteOrder(id string) error {
	return nil
}

func (o *OrderStore) CloseOrder(id string) error {
	return nil
}

func (o *OrderStore) SaveToFile(orders []models.Order) error {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	jsonData, err := json.MarshalIndent(orders, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(o.filePath, jsonData, 0644)
}

func (o *OrderStore) LoadFromFile() error {
	o.mutex.RLock()
	defer o.mutex.RUnlock()

	file, err := os.ReadFile(o.filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) || len(file) == 0 {
			o.orders = []models.Order{}
			o.nextID = 1
			return nil
		}
		return err
	}
	err = json.Unmarshal(file, &o.orders)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) || len(file) == 0 {
			o.orders = []models.Order{}
			o.nextID = 1
			return nil
		}
		return err
	}
	o.nextID = len(o.orders) + 1
	return nil
}
