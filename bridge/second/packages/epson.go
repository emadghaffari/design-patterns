package packages

import "fmt"

// Epson struct
type Epson struct{}

// Print meth
func (e *Epson) Print() {
	fmt.Println("Printing by a EPSON Printer")
}
