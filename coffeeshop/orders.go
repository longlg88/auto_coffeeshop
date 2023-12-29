package coffeeshop

import (
	"auto_coffeeshop/person"
	"auto_coffeeshop/utils"
	"fmt"
	"time"
)

// Orders struct manages the orders in the coffee shop
type Orders struct {
	Recipe   CoffeeRecipe
	Complete chan bool
}

// NewOrder initializes a new coffee order with the given recipe
func NewOrder(recipe CoffeeRecipe) *Orders {
	return &Orders{
		Recipe:   recipe,
		Complete: make(chan bool),
	}
}

func (cs *CoffeeShop) takeOrder(order string, person *person.Person) {
	coffee, exists := cs.Menu[order]

	if !exists {
		fmt.Printf("Sorry, we don't have %s in our menu.\n", order)
		fmt.Println("### End time: ", time.Now().In(utils.Loc).Format("2006-01-02 3:04:05 PM KST"))
		fmt.Println("Waiting...")
		return
	}

	recipe, err := GetCoffeeRecipe(order)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Printf("> %s is received the order now\n", person.Name)

	go func() {
		fmt.Printf("> Starting making coffee for %s\n", coffee.Name)
		time.Sleep(2 * time.Second) // Simulate taking the order
		if err := recipe.MakeCoffee(); err != nil {
			fmt.Println("Error making coffee: ", err)
			cs.Completed <- "Error: Invalid Order"
			return
		}
		cs.Completed <- order
	}()
}
