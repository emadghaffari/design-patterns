package packages

import "fmt"

// ConcreteDecorator struct
type ConcreteDecorator struct {
	De Decorator
}

// Operation method
func (c ConcreteDecorator) Operation() string {
	return fmt.Sprintf("from ConcreteDecorator: %s", c.De.Operation())
}
