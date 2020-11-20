package packages

// CheeseTopping struct
type CheeseTopping struct {
	Pizza pizza
}

// GetPrice meth
func (c *CheeseTopping) GetPrice() int {
	pizzaPrice := c.Pizza.GetPrice()
	return pizzaPrice + 10
}
