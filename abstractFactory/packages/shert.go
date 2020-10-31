package packages

// ShirtInterface interface
type ShirtInterface interface {
    SetLogo(Logo string)
    SetSize(Size int)
    GetLogo() string
    GetSize() int
}

// Shirt interface
type Shirt struct {
    Logo string
    Size int
}

func (s *Shirt) SetLogo(Logo string) {
    s.Logo = Logo
}

func (s *Shirt) GetLogo() string {
    return s.Logo
}

func (s *Shirt) SetSize(Size int) {
    s.Size = Size
}

func (s *Shirt) GetSize() int {
    return s.Size
}