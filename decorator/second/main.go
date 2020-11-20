package main

import (
	"fmt"

	"github.com/emadghaffari/design-patterns/decorator/second/packages"
)

func main() {
	meatPizza := &packages.MeatPizza{}
	fmt.Println("Single Meat Pizza: ", meatPizza.GetPrice())

	meatPizzaWithTomato := &packages.TomatoTopping{Pizza: meatPizza}
	fmt.Println("Meat Pizza with Tomato: ", meatPizzaWithTomato.GetPrice())

	meatPizzaWithTomatoAndCheese := &packages.CheeseTopping{Pizza: meatPizzaWithTomato}
	fmt.Println("Meat Pizza with Tomato and cheese: ", meatPizzaWithTomatoAndCheese.GetPrice())
}
