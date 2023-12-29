package coffeeshop

import (
	"auto_coffeeshop/person"
	"fmt"
	"time"
)

func (cs *CoffeeShop) takeOrder(order string, person *person.Person) {
	if order == "Espresso" {
		fmt.Printf("> %s is received the order now\n", person.Name)
		go func() {
			fmt.Println("> Starting making coffee for Espresso")
			time.Sleep(2 * time.Second) // Simulate taking the order
			cs.makeCoffee(order)
			cs.Completed <- order
		}()
	} else {
		fmt.Println("Sorry, we don't have that in our menu.")
		cs.Completed <- "Error: Invalid Order"
	}
}

func (cs *CoffeeShop) makeCoffee(order string) {
	fmt.Println(">> Grinding coffee beans...")
	time.Sleep(2 * time.Second)

	fmt.Println(">> Putting ground coffee in the filter...")
	time.Sleep(1 * time.Second)

	fmt.Println(">> Extracting espresso...")
	time.Sleep(25 * time.Second)

	fmt.Println(">> Mixing espresso with water...")
	time.Sleep(5 * time.Second)
}
