package packages

import "fmt"

// BarProxy struct
type BarProxy struct {
	Proxbar OpenBar
}

// Welcome method implements Bar interface
func (b *BarProxy) Welcome(h Human) error {
	if h.Age < 20 && h.Gender == "female" {
		return fmt.Errorf("Females under 20 not allowed ")
	}

	if h.Age < 18 && h.Gender == "male" {
		return fmt.Errorf("Males under 18 not allowed ")
	}
	return b.Proxbar.Welcome(h)
}
