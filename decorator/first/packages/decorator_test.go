package packages

import "fmt"

func ExampleDecorator() {
	conComponent := ConcreteComponent{Msg: "method"}
	dec := Decorator{Co: conComponent}

	fmt.Println(dec.Operation())
	// output:
	// from Decorator: method
}
