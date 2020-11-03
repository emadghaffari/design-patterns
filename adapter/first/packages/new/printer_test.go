package new

import (
	"fmt"
	"testing"

	"github.com/Emadghaffari/design-patterns/adapter/packages/old"
	older "github.com/Emadghaffari/design-patterns/adapter/packages/old"
)

func TestPrint(t *testing.T) {
	str := MyPrinter{"Hello World!", older.MyPrinter{}}
	msg := str.Print()
	if msg != "Print: Adapter to New: Hello World!" {
		t.Error("Error in Print for new printer")
	}
}

func ExamplePrint() {
	str := MyPrinter{"Hello World!", old.MyPrinter{}}
	msg := str.Print()
	fmt.Println(msg)
	// Output:
	// Print: Adapter to New: Hello World!
}
