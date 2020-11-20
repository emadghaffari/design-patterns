package packages

import "fmt"

// Linux struct
type Linux struct {
	printer printer
}

// SetPrinter meth
func (m *Linux) SetPrinter(printer printer) {
	m.printer = printer
}

// Print meth
func (m *Linux) Print() {
	fmt.Println("Print request for Linux")
	m.printer.Print()
}
