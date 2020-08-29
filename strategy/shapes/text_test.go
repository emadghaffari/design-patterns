package shapes

import (
	"os"
	"testing"
)

func TestTextPrint(t *testing.T) {
	sq := TextSquare{}
	sq.SetWriter(os.Stdout)
	sq.Print()
}

// ExamplePrint func
func ExampleTextPrint() {
	sq := TextSquare{}
	sq.SetWriter(os.Stdout)
	sq.Print()
	// Output:
	// circle
}
