package packages

import "fmt"

func ExampleConcreteComponent() {
	conComponent := ConcreteComponent{Msg: "ConcreteComponent from Operation method"}

	fmt.Println(conComponent.Operation())
	// Output:
	// ConcreteComponent from Operation method
}
