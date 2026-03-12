package main

import (
	"fmt"

	"inventory/internal/models"
	"inventory/internal/services"
)

func main() {
	fmt.Println("Inventory System")

	inventory := services.NewInventory()
	itens := [] models.Item {
		{ID: 1, Name: "Phone", Quantity: 3, Price: 79.99},
		{ID: 2, Name: "T-Shirt", Quantity: 7, Price: 39.99},
		{ID: 3, Name: "Mouse", Quantity: 2, Price: 19.99},	
	}

	for _, item := range itens {
		err := inventory.AddItem(item, "Alurete")
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	listItens := inventory.ListItems()
	fmt.Printf("\n ::: TOTAL ITENS IN INVENTORY: %d :::\n", len(listItens))
	for _, item := range listItens {
		fmt.Printf("\n -- %s", item.Info())
	}

	// fmt.Println("\n Total Inventory $", inventory.CalculateTotalCost())

	// logs := inventory.ViewAuditLog()
	// for _, log := range logs {
	// 	fmt.Printf("\n -- %s", log.Info())
	// }

	// itemToSearch, err := services.FindBy(itens, func(item models.Item) bool {
	// 	return item.Price > 29.99
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("Founded: ", itemToSearch)

	// alura := services.Supplier {
	// 	CNPJ: "123456",
	// 	Contact: "15997854466",
	// 	City: "Sorocaba",
	// }

	// fmt.Println(alura.GetInfo())
	// if alura.VerifyAvailability(100, 15) {
	// 	fmt.Println("It has availability!! XD")
	// } else {
	// 	fmt.Println("It has *NO* availability!! :(")
	// }

	fmt.Println("\n\n::: XXXXX :::")

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

	fmt.Println("=====")

	listItens = inventory.ListItems()
	for _, item := range listItens {
		fmt.Printf("\n -- %s", item.Info())
	}
	
	err := inventory.RemoveItem(1, 1, "Alurete")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("")
	fmt.Println(".....")
	logs := inventory.ViewAuditLog()
	for _, log := range logs {
		fmt.Printf("\n -- %s", log.Info())
	}
	fmt.Println("")
	fmt.Println(".....")

	listItens = inventory.ListItems()
	for _, item := range listItens {
		fmt.Printf("\n -- %s", item.Info())
	}
	fmt.Println("")

	listItens = inventory.ListItems()
	itemByName, err := services.FindByName(listItens, "T-Shirt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\n+++++")
	fmt.Println(itemByName)

	listItens = inventory.ListItems()
	itemBy, err := services.FindBy(listItens, func(item models.Item) bool {
		return item.Price > 20.00
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\n*****")
	fmt.Println(itemBy)
}
