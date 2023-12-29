package main

import (
	"time"

	"auto_coffeeshop/coffeeshop"
)

func main() {
	coffeeShop := coffeeshop.NewCoffeeShop()

	go coffeeShop.StartService()

	// Simulate order
	// You can continue adding orders here as needed
	// For an infinite loop, you can keep adding orders
	for {
		coffeeShop.Orders <- "Espresso"
		time.Sleep(1 * time.Minute) // Wait before placing another order
		// select {
		// case <-coffeeShop.Completed:
		// 	fmt.Println("Waiting...")
		// 	time.Sleep(1 * time.Minute) // Wait before placing another order
		// }
	}
}
