package packages

// ShoeInterface interface
type ShoeInterface interface{
	SetLogo(Logo string)
    SetSize(Size int)
    GetLogo() string
    GetSize() int
}

// Shoe struct
type Shoe struct {
    Logo string
    Size int
}

func (s *Shoe) SetLogo(Logo string) {
    s.Logo = Logo
}

func (s *Shoe) GetLogo() string {
    return s.Logo
}

func (s *Shoe) SetSize(Size int) {
    s.Size = Size
}

func (s *Shoe) GetSize() int {
    return s.Size
}