package packages

// Adidas struct
type Adidas struct{}

type AdidasShoe struct {
	Shoe
}
type AdidasShirt struct {
	Shirt
}

// MakeShoe meth
func (a *Adidas) MakeShoe() ShoeInterface {
    return &AdidasShoe{
        Shoe: Shoe{
            Logo: "adidas",
            Size: 14,
        },
    }
}

// MakeShirt meth
func (a *Adidas) MakeShirt() ShirtInterface {
    return &AdidasShirt{
        Shirt: Shirt{
            Logo: "adidas",
            Size: 14,
        },
    }
}