package new

import (
	"fmt"

	"github.com/Emadghaffari/design-patterns/adapter/packages/old"
)

type printer interface {
	Print() string
}

// MyPrinter struct
type MyPrinter struct {
	Msg string
	old.MyPrinter
}

// Print method MyPrinter
func (p *MyPrinter) Print() (new string) {
	new = p.MyPrinter.Print(fmt.Sprintf("Adapter to New: %s", p.Msg))
	return
}
