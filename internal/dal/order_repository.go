package dal

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/IceMAN2377/hot-coffee/internal/models"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

type OrderRepository interface {
	CreateOrder(ord *models.CreateOrderMod) (*models.Order, error)
	GetAll() ([]models.Order, error)
	GetOrder(id string) (*models.Order, error)
	UpdateOrder(id string, items []models.OrderItem) (*models.Order, error)
	DeleteOrder(id string) error
	CloseOrder(id string) error
	CalculateIngredients(items []models.OrderItem) (map[string]models.InventoryItem, error)
}

type OrderStore struct {
	//logger   slog.Logger
	filePath string
	mutex    sync.RWMutex
	orders   []models.Order
	nextID   int
	menuRepo MenuRepository
	invRepo  InventoryRepository
}

func NewOrderStore(filePath string, menuRepo MenuRepository, invRepo InventoryRepository) *OrderStore {
	o := &OrderStore{
		//logger:   *logger,
		filePath: filePath,
		nextID:   1,
		menuRepo: menuRepo,
		invRepo:  invRepo,
	}
	err := o.LoadFromFile()
	if err != nil {
		log.Printf("error loading: %v", err)
	}
	return o
}

func (o *OrderStore) CreateOrder(ord *models.CreateOrderMod) (*models.Order, error) {
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

	requiredIngredients, err := o.CalculateIngredients(ord.Items)
	if err != nil {
		return nil, err
	}

	isEnough, err := o.invRepo.CheckAvailability(requiredIngredients)
	if err != nil {
		return nil, err
	}

	if !isEnough {
		return nil, errors.New("insufficient ingredients for order")
	}

	if err := o.invRepo.DeductFromInventory(requiredIngredients); err != nil {
		return nil, errors.New("error deducting from inventory")
	}

	o.orders = append(o.orders, order)

	if err := o.SaveToFile(o.orders); err != nil {
		log.Fatal("failed to save to order json")
	}
	o.nextID++

	return &order, nil
}

func (o *OrderStore) GetAll() ([]models.Order, error) {
	return o.orders, nil
}

func (o *OrderStore) GetOrder(id string) (*models.Order, error) {
	idInt, _ := strconv.Atoi(id)
	err := o.LoadFromFile()
	if err != nil {
		log.Printf("error loading by id :%v", err)
	}
	return &o.orders[idInt-1], nil
}

func (o *OrderStore) UpdateOrder(id string, items []models.OrderItem) (*models.Order, error) {

	idInt, _ := strconv.Atoi(id)

	err := o.LoadFromFile()
	if err != nil {
		log.Printf("error loading by id :%v", err)
	}

	if len(o.orders) < idInt || len(o.orders) > idInt {
		log.Printf("id doesnot exist: %v", err)
		return nil, err
	}

	o.orders[idInt-1].Items = items

	if err := o.SaveToFile(o.orders); err != nil {
		log.Fatal("failed to save to order json")
	}

	return &o.orders[idInt-1], nil
}

func (o *OrderStore) DeleteOrder(id string) error {
	idInt, _ := strconv.Atoi(id)

	err := o.LoadFromFile()
	if err != nil {
		log.Printf("error loading by id :%v", err)
	}

	o.orders[idInt-1] = models.Order{}

	if err := o.SaveToFile(o.orders); err != nil {
		log.Fatal("failed to save to order json")
	}

	return nil
}

func (o *OrderStore) CloseOrder(id string) error {
	idInt, _ := strconv.Atoi(id)

	err := o.LoadFromFile()
	if err != nil {
		log.Printf("error loading by id :%v", err)
	}

	o.orders[idInt-1].Status = "closed"

	if err := o.SaveToFile(o.orders); err != nil {
		log.Fatal("failed to save to order json")
	}

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

func (o *OrderStore) CalculateIngredients(items []models.OrderItem) (map[string]models.InventoryItem, error) {
	reqIngr := make(map[string]models.InventoryItem)

	for _, item := range items {
		menuPosition, err := o.menuRepo.GetItem(item.ProductID)
		if err != nil {
			return nil, err
		}
		for _, menuItemIngred := range menuPosition.Ingredients {
			required := menuItemIngred.Quantity * float64(item.Quantity)

			if existing, exists := reqIngr[menuItemIngred.IngredientID]; exists {
				// Aggregate quantities for same ingredient
				existing.Quantity += required
				reqIngr[menuItemIngred.IngredientID] = existing
			} else {
				// Create new inventory item for this ingredient
				reqIngr[menuItemIngred.IngredientID] = models.InventoryItem{
					IngredientID: menuItemIngred.IngredientID,
					Name:         "", // Will be populated when checking inventory
					Quantity:     required,
					Unit:         "",
				}
			}
		}
	}
	return reqIngr, nil
}
