package packages

// TomatoTopping struct
type TomatoTopping struct {
	Pizza pizza
}

// GetPrice meth
func (c *TomatoTopping) GetPrice() int {
	pizzaPrice := c.Pizza.GetPrice()
	return pizzaPrice + 7
}
