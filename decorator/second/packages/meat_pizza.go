package packages

// MeatPizza struct
type MeatPizza struct {
	Price int
}

// GetPrice meth for MeatPizza
func (p *MeatPizza) GetPrice() int {
	return 15
}
