package main

import (
	"fmt"

	"github.com/Emadghaffari/design-patterns/decorator/packages"
)

func ExampleConcreteComponent() {
	conComponent := packages.ConcreteComponent{Msg: "ConcreteComponent from Operation method"}

	fmt.Println(conComponent.Operation())
	// Output:
	// ConcreteComponent from Operation method
}

func ExampleDecorator() {
	conComponent := packages.ConcreteComponent{Msg: "ConcreteComponent from Operation method"}
	dec := packages.Decorator{Co: conComponent}

	fmt.Println(dec.Operation())
	// output:
	// from Decorator: ConcreteComponent from Operation method
}

func ExampleConcreteDecorator() {
	conComponent := packages.ConcreteComponent{Msg: "ConcreteComponent from Operation method"}
	dec := packages.Decorator{Co: conComponent}
	conDecorator := packages.ConcreteDecorator{De: dec}

	fmt.Println(conDecorator.Operation())
	// Output:
	// from ConcreteDecorator: from Decorator: ConcreteComponent from Operation method
}
