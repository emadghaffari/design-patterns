package factory

import (
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/emadghaffari/design-patterns/strategy/shapes"
)

func TestNewPrint(t *testing.T) {
	i, err := NewPrint(TextStrategy)
	if err != nil {
		t.Error(err.Error())
	}
	if reflect.TypeOf(i) != reflect.TypeOf(&shapes.TextSquare{}) {
		t.Error("the return type for TextStrategy is not ok!")
	}

	j, err := NewPrint(ImageStrategy)
	if err != nil {
		t.Error(err.Error())
	}
	if reflect.TypeOf(j) != reflect.TypeOf(&shapes.ImageSquare{}) {
		t.Error("the return type for ImageStrategy is not ok!")
	}

	k, err := NewPrint(PdfStrategy)
	if err != nil {
		t.Error(err.Error())
	}
	if reflect.TypeOf(k) != reflect.TypeOf(&shapes.PdfSquare{}) {
		t.Error("the return type for PdfStrategy is not ok!")
	}
}

func ExampleNewPrint() {
	i, err := NewPrint(TextStrategy)
	if err != nil {
		log.Fatalln(err.Error())
	}
	if reflect.TypeOf(i) != reflect.TypeOf(&shapes.TextSquare{}) {
		log.Fatalln("the return type for TextStrategy is not ok!")
	}

	fmt.Print(reflect.TypeOf(i))
	// Output:
	// *shapes.TextSquare

}
