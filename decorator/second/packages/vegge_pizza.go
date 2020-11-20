package packages

// VeggePizza struct
type VeggePizza struct {
	Price int
}

// GetPrice meth for VeggePizza
func (p *VeggePizza) GetPrice() int {
	return 7
}
