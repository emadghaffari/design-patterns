package packages

// Nike struct
type Nike struct {
}

type NikeShoe struct{
	Shoe
}

type NikeShirt struct{
	Shirt
}

// MakeShoe meth
func (n *Nike) MakeShoe() ShoeInterface {
    return &NikeShoe{
        Shoe: Shoe{
            Logo: "nike",
            Size: 14,
        },
    }
}

// MakeShirt meth
func (n *Nike) MakeShirt() ShirtInterface {
    return &NikeShirt{
        Shirt: Shirt{
            Logo: "nike",
            Size: 14,
        },
    }
}