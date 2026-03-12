package services

import (
	"fmt"
	"strconv"
	"time"
	"inventory/internal/models"
)

type Inventory struct {
	items map[string] models.Item
	logs [] models.Log
}

func NewInventory() * Inventory {
	return &Inventory {
		items: make(map[string] models.Item),
		logs: [] models.Log {},
	}
}

func (i *Inventory) AddItem(item models.Item, user string) error {
	if item.Quantity <= 0 {
		return fmt.Errorf("Error to add item: [ID: %d] - invalid quantity (less or equal than zero)", item.ID)
	}

	existingItem, exists := i.items[strconv.Itoa(item.ID)]
	if exists {
		item.Quantity += existingItem.Quantity
	}
	i.items[strconv.Itoa(item.ID)] = item

	i.logs = append(i.logs, models.Log {
		Timestamp: time.Now(),
		Action: "Add Item in Inventory",
		User: user,
		ItemID: item.ID,
		Quantity: item.Quantity,
		Reason: "New Item must be added into inventory",
	})

	return nil
}

func (i *Inventory) ListItems() [] models.Item {
	var itemList [] models.Item

	for _, item := range i.items {
		itemList = append(itemList, item)
	}

	return itemList
}

func (i *Inventory) ViewAuditLog() [] models.Log {
	return i.logs
}

func (i *Inventory) CalculateTotalCost() float64 {
	var totalCost float64

	for _, item := range i.items {
		totalCost += float64(item.Quantity) * item.Price
	}

	return totalCost
}

func FindBy[T any](data []T, comparator func(T) bool) ([] T, error) {
	var result [] T

	for _, v := range data {
		if comparator(v) {
			result = append(result, v)
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("Item not found!!")
	}

	return result, nil
}

func FindByName(data [] models.Item, name string) ([] models.Item, error) {
	var result [] models.Item

	for _, item := range data {
		if item.Name == name {
			result = append(result, item)
		}

	}
	if len(result) == 0 {
		return nil, fmt.Errorf("No item with name %s was found!!", name)
	}

	return result, nil
}

func (i *Inventory) RemoveItem(itemID int, quantity int, user string) error {
	existingItem, exists := i.items[strconv.Itoa(itemID)]
	if !exists {
		return fmt.Errorf("Error to remove item: [ID: %d] do not exists!!", itemID)
	}

	if quantity <= 0 {
		return fmt.Errorf("Error to remove item: [ID: %d] less or equal to zero (%d)!!", itemID, quantity)
	}

	if existingItem.Quantity < quantity {
		return fmt.Errorf("Error to remove item: [ID: %d] existing item quantity (%d) is insufficient (%d)!!", itemID, existingItem.Quantity, quantity)
	}

	existingItem.Quantity -= quantity
	if existingItem.Quantity == 0 {
		delete(i.items, strconv.Itoa(itemID))
	} else {
		i.items[strconv.Itoa(itemID)] = existingItem
	}

	i.logs = append(i.logs, models.Log {
		Timestamp: time.Now(),
		Action: "Remove Item in Inventory",
		User: user,
		ItemID: itemID,
		Quantity: quantity,
		Reason: "REMOVED Item from inventory",
	})

	return nil
}
