package packages

// ConcreteComponent struct
type ConcreteComponent struct {
	Msg string
}

// Operation method
func (c ConcreteComponent) Operation() string {
	return c.Msg
}
