package main

import (
	"fmt"
	"log"

	"github.com/Emadghaffari/design-patterns/proxy/packages"
)

func main() {
	h1 := packages.Human{
		Name:   "jon",
		Age:    12,
		Gender: "male",
	}
	h2 := packages.Human{
		Name:   "sozan",
		Age:    15,
		Gender: "female",
	}
	bar := packages.OpenBar{}
	// jon is 12 !?
	if err := bar.Welcome(h1); err != nil {
		log.Fatalln("Error in welcome: ", err.Error())
	}
	// sozan 15 ?!
	if err := bar.Welcome(h2); err != nil {
		log.Fatalln("Error in welcome: ", err.Error())
	}
	fmt.Println(bar)

	fmt.Println("---------------------proxy----------------------")

	h3 := packages.Human{
		Name:   "jon",
		Age:    12,
		Gender: "male",
	}
	h4 := packages.Human{
		Name:   "sozan",
		Age:    15,
		Gender: "female",
	}
	h5 := packages.Human{
		Name:   "sozan",
		Age:    55,
		Gender: "female",
	}
	proxyBar := packages.BarProxy{}
	if err := proxyBar.Welcome(h3); err != nil {
		fmt.Println("Error in welcome: ", err.Error())
	}
	// sozan 15 ?!
	if err := proxyBar.Welcome(h4); err != nil {
		fmt.Println("Error in welcome: ", err.Error())
	}
	if err := proxyBar.Welcome(h5); err != nil {
		fmt.Println("Error in welcome: ", err.Error())
	}
	fmt.Println(proxyBar)
}
