package packages

import (
	"fmt"
)

// Decorator struct
type Decorator struct {
	Co ConcreteComponent
}

// Operation method
func (de Decorator) Operation() string {
	return fmt.Sprintf("from Decorator: %s", de.Co.Operation())
}
