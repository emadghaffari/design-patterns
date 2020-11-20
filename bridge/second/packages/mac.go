package packages

import "fmt"

// Mac struct
type Mac struct {
	printer printer
}

// SetPrinter meth
func (m *Mac) SetPrinter(printer printer) {
	m.printer = printer
}

// Print meth
func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.Print()
}
