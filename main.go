package main

import (
	"math/rand"
	"time"

	"auto_coffeeshop/coffeeshop"
)

func main() {
	coffeeShop := coffeeshop.NewCoffeeShop()

	go coffeeShop.StartService()

	// Simulate order with random coffee types
	coffeeTypes := []string{"Espresso", "Latte", "Cappuccino"} // Add more coffee types as needed

	for {
		// Generate a random index within the coffeeTypes slice length
		randomIndex := rand.Intn(len(coffeeTypes))

		// Select a random coffee type using the generated index
		selectedCoffee := coffeeTypes[randomIndex]

		coffeeShop.Orders <- selectedCoffee
		time.Sleep(1 * time.Minute) // Wait before placing another order
	}
}
