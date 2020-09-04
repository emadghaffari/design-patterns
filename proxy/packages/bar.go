package packages

// Bar interface
type Bar interface {
	Welcome(Human) error
}

// OpenBar struct
type OpenBar struct {
	Visitors []Human
}

// Welcome method implements Bar interface
func (b *OpenBar) Welcome(h Human) error {
	b.Visitors = append(b.Visitors, h)
	return nil
}
