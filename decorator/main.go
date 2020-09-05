package main

import (
	"fmt"

	"github.com/Emadghaffari/design-patterns/decorator/packages"
)

func main() {
	conComponent := packages.ConcreteComponent{Msg: "ConcreteComponent from Operation method"}
	fmt.Println(conComponent.Operation())

	dec := packages.Decorator{Co: conComponent}
	fmt.Println(dec.Operation())

	conDecorator := packages.ConcreteDecorator{De: dec}
	fmt.Println(conDecorator.Operation())

}
