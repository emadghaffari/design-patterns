package packages

import (
	"fmt"
)

// tshirt colors
const (
	WHITE  = 1
	GREEN  = 2
	BLACK  = 3
	YELLOW = 4
)

// ProductInfoGetter interface
type ProductInfoGetter interface {
	GetInfo() string
}

// Tshirt struct
type Tshirt struct {
	Size  string
	Color string
}

// GetInfo method
func (t Tshirt) GetInfo() string {
	return fmt.Sprintf("Tshirt size: '%s', Color: '%s'.", t.Size, t.Color)
}

var whiteTshirt *Tshirt = &Tshirt{"M", "white"}
var greenTshirt *Tshirt = &Tshirt{"L", "green"}
var blackTshirt *Tshirt = &Tshirt{"S", "black"}
var yellowTshirt *Tshirt = &Tshirt{"XS", "yellow"}

// GetTshirClone func
func GetTshirClone(t int) (ProductInfoGetter, error) {
	switch t {
	case WHITE:
		return whiteTshirt, nil
	case GREEN:
		return greenTshirt, nil
	case BLACK:
		return blackTshirt, nil
	case YELLOW:
		return yellowTshirt, nil
	default:
		return nil, fmt.Errorf(fmt.Sprintf("Prototype id %d not recognized.", t))
	}
}
