package coffeeshop

import (
	"fmt"
	"time"

	"auto_coffeeshop/person"
)

type Coffee struct {
	Name        string
	Price       float64
	Ingredients []string
}

type CoffeeShop struct {
	Menu      map[string]Coffee
	Orders    chan string
	Completed chan string
}

func NewCoffeeShop() *CoffeeShop {
	menu := map[string]Coffee{
		"Espresso": {Name: "Espresso", Price: 1.5, Ingredients: []string{"Coffee"}},
		"Latte":    {Name: "Latte", Price: 3.0, Ingredients: []string{"Coffee", "Milk"}},
		// Add more coffee types here
	}

	return &CoffeeShop{
		Menu:      menu,
		Orders:    make(chan string),
		Completed: make(chan string),
	}
}

func (cs *CoffeeShop) StartService() {
	loc, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	employee := &person.Person{Name: "Reacher"}

	for {
		select {
		case order := <-cs.Orders:
			fmt.Println("### Start time: ", time.Now().In(loc).Format("2006-01-02 3:04:05 PM KST"))
			cs.takeOrder(order, employee)
		case completedOrder := <-cs.Completed:
			fmt.Printf("> Completed order. Here is your %s! Enjoy!\n", completedOrder)
			fmt.Println("### End time: ", time.Now().In(loc).Format("2006-01-02 3:04:05 PM KST"))
			fmt.Println("Waiting...")
		}
	}
}
