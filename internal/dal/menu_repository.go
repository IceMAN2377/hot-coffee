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

type MenuRepository interface {
	AddItem(item *models.MenuItem) (*models.MenuItem, error)
	GetItems() ([]models.MenuItem, error)
	GetItem(id string) (*models.MenuItem, error)
	UpdateItem(id string, item *models.MenuItem) (*models.MenuItem, error)
	DeleteItem(id string) error
}

type MenuStore struct {
	filePath  string
	mutex     sync.RWMutex
	menuItems []models.MenuItem
}

func NewMenuStore(filePath string) *MenuStore {
	m := &MenuStore{
		filePath: filePath,
	}

	err := m.LoadFromFile()
	if err != nil {
		log.Printf("error loading: %v", err)
	}
	return m
}

func (m *MenuStore) AddItem(item *models.MenuItem) (*models.MenuItem, error) {
	id := strings.ReplaceAll(strings.ToLower(item.Name), " ", "_")

	item.ID = id
	m.menuItems = append(m.menuItems, *item)

	if err := m.SaveToFile(m.menuItems); err != nil {
		log.Fatal("failed to save to order json")
	}

	return item, nil
}

func (m *MenuStore) GetItems() ([]models.MenuItem, error) {
	return m.menuItems, nil
}

func (m *MenuStore) GetItem(id string) (*models.MenuItem, error) {

	for _, items := range m.menuItems {
		if id == items.ID {
			return &items, nil
		}
	}
	return &models.MenuItem{}, nil
}

func (m *MenuStore) UpdateItem(id string, item *models.MenuItem) (*models.MenuItem, error) {
	found := false
	var updatedItem *models.MenuItem

	for i := 0; i < len(m.menuItems); i++ {
		if id == m.menuItems[i].ID {
			found = true
			if item.Name != "" {
				m.menuItems[i].Name = item.Name
				newID := strings.ReplaceAll(strings.ToLower(item.Name), " ", "_")
				m.menuItems[i].ID = newID
			}
			if item.Description != "" {
				m.menuItems[i].Description = item.Description
			}
			if item.Price != 0 {
				m.menuItems[i].Price = item.Price
			}
			if item.Ingredients != nil {
				m.menuItems[i].Ingredients = item.Ingredients
			}
			updatedItem = &m.menuItems[i]
			break
		}
	}
	if !found {
		return nil, errors.New("item not found")
	}

	if err := m.SaveToFile(m.menuItems); err != nil {
		log.Fatal("failed to save to menu json")
	}

	return updatedItem, nil
}

func (m *MenuStore) DeleteItem(id string) error {
	return nil
}

func (m *MenuStore) SaveToFile(menu []models.MenuItem) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	jsonData, err := json.MarshalIndent(menu, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(m.filePath, jsonData, 0644)
}

func (m *MenuStore) LoadFromFile() error {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	file, err := os.ReadFile(m.filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) || len(file) == 0 {
			m.menuItems = []models.MenuItem{}
			return nil
		}
		return err
	}
	err = json.Unmarshal(file, &m.menuItems)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) || len(file) == 0 {
			m.menuItems = []models.MenuItem{}
			return nil
		}
		return err
	}
	return nil
}
