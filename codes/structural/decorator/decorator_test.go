package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	veggePizza := &veggeMania{}

	// add cheese topping
	veggePizzaWithCheese := &cheeseTopping{
		pizza: veggePizza,
	}

	// add tomato topping
	veggePizzaWithCheeseAndTomato := &tomatoTopping{
		pizza: veggePizzaWithCheese,
	}

	fmt.Printf("Price of veggieMania pizza with tomato and cheese topping is %d\n", veggePizzaWithCheeseAndTomato.getPrice())

	peppyPannerPrizza := &peppyPaneer{}

	// add cheese topping
	peppyPannerPrizzaWithCheese := &cheeseTopping{
		pizza: peppyPannerPrizza,
	}

	fmt.Printf("Price of peppyPaneer with tomato and cheese topping is %d\n", peppyPannerPrizzaWithCheese.getPrice())
}
