package packages

import "fmt"

func ExampleConcreteDecorator() {
	conComponent := ConcreteComponent{Msg: "ConcreteComponent"}
	dec := Decorator{Co: conComponent}
	conDecorator := ConcreteDecorator{De: dec}

	fmt.Println(conDecorator.Operation())
	// Output:
	// from ConcreteDecorator: from Decorator: ConcreteComponent
}
