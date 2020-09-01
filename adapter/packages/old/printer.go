package old

import "fmt"

type printer interface {
	Print(s string) string
}

// MyPrinter struct
type MyPrinter struct{}

// Print method MyPrinter
func (p *MyPrinter) Print(s string) (new string) {
	new = fmt.Sprintf("Print: %s", s)
	return
}
