package dal

import (
	"encoding/json"
	"errors"
	"github.com/IceMAN2377/hot-coffee/internal/models"
	"log"
	"os"
	"strings"
	"sync"
)

type InventoryRepository interface {
	AddItem(item *models.InventoryItem) (*models.InventoryItem, error)
	GetItems() ([]models.InventoryItem, error)
	GetItem(id string) (*models.InventoryItem, error)
	UpdateItem(id string, item *models.InventoryItem) (*models.InventoryItem, error)
	DeleteItem(id string) error
}

type InventoryStore struct {
	filePath string
	mutex    sync.RWMutex
	invItems []models.InventoryItem
}

func NewInventoryStore(filePath string) *InventoryStore {
	inv := &InventoryStore{
		filePath: filePath,
	}

	if err := inv.LoadFromFile(); err != nil {
		log.Printf("error loading from file: %v", err)
	}

	return inv
}

func (v *InventoryStore) AddItem(item *models.InventoryItem) (*models.InventoryItem, error) {
	id := strings.ReplaceAll(strings.ToLower(item.Name), " ", "_")

	item.IngredientID = id
	v.invItems = append(v.invItems, *item)

	if err := v.SaveToFile(v.invItems); err != nil {
		log.Fatal("failed to save to order json")
	}

	return item, nil
}

func (v *InventoryStore) GetItems() ([]models.InventoryItem, error) {
	return v.invItems, nil
}

func (v *InventoryStore) GetItem(id string) (*models.InventoryItem, error) {

	for _, items := range v.invItems {
		if id == items.IngredientID {
			return &items, nil
		}
	}
	return nil, errors.New("item not found")
}

func (v *InventoryStore) UpdateItem(id string, item *models.InventoryItem) (*models.InventoryItem, error) {
	found := false
	var updatedItem *models.InventoryItem

	for i := 0; i < len(v.invItems); i++ {
		if id == v.invItems[i].IngredientID {
			found = true
			if item.Name != "" {
				v.invItems[i].Name = item.Name
				newID := strings.ReplaceAll(strings.ToLower(item.Name), " ", "_")
				v.invItems[i].IngredientID = newID
			}
			if item.Quantity != 0 {
				v.invItems[i].Quantity = item.Quantity
			}
			if item.Unit != "" {
				v.invItems[i].Unit = item.Unit
			}

			updatedItem = &v.invItems[i]
			break
		}
	}
	if !found {
		return nil, errors.New("item not found")
	}

	if err := v.SaveToFile(v.invItems); err != nil {
		log.Fatal("failed to save to menu json")
	}

	return updatedItem, nil
}

func (v *InventoryStore) DeleteItem(id string) error {

	for i := 0; i < len(v.invItems); i++ {
		if id == v.invItems[i].IngredientID {
			v.invItems = append(v.invItems[:i], v.invItems[i+1:]...)
			return v.SaveToFile(v.invItems)
		}
	}

	return errors.New("item not found")
}

func (v *InventoryStore) SaveToFile(inv []models.InventoryItem) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	jsonData, err := json.MarshalIndent(inv, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(v.filePath, jsonData, 0644)
}

func (v *InventoryStore) LoadFromFile() error {
	v.mutex.RLock()
	defer v.mutex.RUnlock()

	file, err := os.ReadFile(v.filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) || len(file) == 0 {
			v.invItems = []models.InventoryItem{}
			return nil
		}
		return err
	}
	err = json.Unmarshal(file, &v.invItems)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) || len(file) == 0 {
			v.invItems = []models.InventoryItem{}
			return nil
		}
		return err
	}
	return nil
}
