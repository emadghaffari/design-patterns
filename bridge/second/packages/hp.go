package packages

import "fmt"

// Hp struct
type Hp struct{}

// Print meth
func (e *Hp) Print() {
	fmt.Println("Printing by a Hp Printer")
}
