package packages

import (
	"fmt"
)

func ExampleHuman() {
	h1 := Human{
		"Emad",
		25,
		"male",
	}
	fmt.Println(h1)
	// Output:
	// {Emad 25 male}
}
