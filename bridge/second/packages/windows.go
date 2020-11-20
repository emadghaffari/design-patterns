package packages

import "fmt"

// Windows struct
type Windows struct {
	printer printer
}

// SetPrinter meth
func (m *Windows) SetPrinter(printer printer) {
	m.printer = printer
}

// Print meth
func (m *Windows) Print() {
	fmt.Println("Print request for Windows")
	m.printer.Print()
}
