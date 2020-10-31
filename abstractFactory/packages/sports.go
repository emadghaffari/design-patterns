package packages

import "fmt"

// SportsFactory interface
// this interface implement shoe and shirt to another classes
type SportsFactory interface {
    MakeShoe() ShoeInterface
	MakeShirt() ShirtInterface
}


// GetSportsFactory func
// for get a sport factory by name
// example adidas, nike, ....
func GetSportsFactory(brand string) (SportsFactory, error) {
	if brand == ""{
		return nil, fmt.Errorf("Error in set brand")
	}

	switch brand {
	case "adidas":
		return &Adidas{}, nil
	case "nike":
		return &Nike{}, nil
	default:
		return nil,fmt.Errorf("brand not found")
	}
}