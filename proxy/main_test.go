package main

import (
	"fmt"
	"log"

	"github.com/Emadghaffari/design-patterns/proxy/packages"
)

func ExampleBar() {
	h1 := packages.Human{
		Name:   "jon",
		Age:    12,
		Gender: "male",
	}
	bar := packages.OpenBar{}
	// jon is 12 !?
	if err := bar.Welcome(h1); err != nil {
		log.Fatalln("Error in welcome: ", err.Error())
	}
	fmt.Println(bar)
	// Output:
	// {[{jon 12 male}]}
}

func ExampleProxyBar() {
	h1 := packages.Human{
		Name:   "jon",
		Age:    12,
		Gender: "male",
	}
	proxyBar := packages.BarProxy{}
	if err := proxyBar.Welcome(h1); err != nil {
		fmt.Println("Error in welcome: ", err.Error())
	}
	// Output:
	// Error in welcome:  Males under 18 not allowed
}
