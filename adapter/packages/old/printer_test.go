package old

import (
	"fmt"
	"testing"
)

func TestPrint(t *testing.T) {
	str := MyPrinter{}
	msg := str.Print("Hello World!")

	if msg != "Print: Hello World!" {
		t.Error("Error in Print for old printer")
	}
}

func ExamplePrint() {
	str := MyPrinter{}
	msg := str.Print("Hello World!")
	fmt.Println(msg)
	// Output:
	// Print: Hello World!
}
