package main

import (
	"fmt"

	"inventory/internal/models"
)

func main() {
	fmt.Println("Inventory System")

	item1 := models.Item {
		ID: 1,
		Name: "First Product",
		Quantity: 10,
		Price: 9.99,
	}

	item2 := models.Item {
		ID: 2,
		Name: "Second Product",
		Quantity: 20,
		Price: 99.99,
	}

	fmt.Println(item1)
	fmt.Println(item1.ID)
	fmt.Println(item1.Info())

	fmt.Println("-----")
	
	fmt.Println(item2)
	fmt.Println(item2.ID)
	fmt.Println(item2.Info())
}
