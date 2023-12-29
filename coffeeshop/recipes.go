package coffeeshop

import (
	"errors"
	"fmt"
	"time"
)

// CoffeeRecipe interface defines the method for making different coffees
type CoffeeRecipe interface {
	MakeCoffee() error
}

// EspressoRecipe represents the recipe for making an Espresso
type EspressoRecipe struct{}

// MakeCoffee implements the CoffeeRecipe interface for Espresso
func (e *EspressoRecipe) MakeCoffee() error {
	// Recipe steps to make Espresso
	fmt.Println(">> Grinding coffee beans...")
	time.Sleep(2 * time.Second)

	fmt.Println(">> Putting ground coffee in the filter...")
	time.Sleep(1 * time.Second)

	fmt.Println(">> Extracting espresso...")
	time.Sleep(25 * time.Second)

	fmt.Println(">> Mixing espresso with water...")
	time.Sleep(5 * time.Second)

	return nil
}

// LatteRecipe represents the recipe for making a Latte
type LatteRecipe struct{}

// MakeCoffee implements the CoffeeRecipe interface for Latte
func (l *LatteRecipe) MakeCoffee() error {
	// Recipe steps to make Latte
	fmt.Println(">> Grinding coffee beans...")
	time.Sleep(2 * time.Second)

	fmt.Println(">> Frothing milk...")
	time.Sleep(1 * time.Second)

	fmt.Println(">> Extracting espresso...")
	time.Sleep(25 * time.Second)

	fmt.Println(">> Mixing espresso with milk...")
	time.Sleep(5 * time.Second)

	return nil
}

// GetCoffeeRecipe returns the appropriate CoffeeRecipe based on the coffee type
func GetCoffeeRecipe(coffeeType string) (CoffeeRecipe, error) {
	switch coffeeType {
	case "Espresso":
		return &EspressoRecipe{}, nil
	case "Latte":
		return &LatteRecipe{}, nil
	default:
		return nil, errors.New("recipe not found for the given coffee type")
	}
}
