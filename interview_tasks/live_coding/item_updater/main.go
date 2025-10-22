package main

import "fmt"

type Item struct {
	Name  *string
	Price int
}

func updateItem(item *Item, newName string, newPrice int) {
	item.Name = &newName
	item.Price = newPrice
	//fmt.Println(*item.Name, item.Price) // Книга в мягкой обложке 10
}

func main() {
	defaultName := "Книга"
	item := &Item{
		Name:  &defaultName, // 0x000...
		Price: 15,           // 15
	}
	//fmt.Println(*item.Name, item.Price) // Книга 15

	secondaryItem := Item{} // *item // Книга 15

	//fmt.Println(*secondaryItem.Name, secondaryItem.Price)

	updateItem(&secondaryItem, "Книга в мягкой обложке", 10)

	fmt.Println(*secondaryItem.Name, secondaryItem.Price) // &Item{0x001..., 10}
	fmt.Println(*item.Name, item.Price)                   // &Item{0x001..., 10}
}
