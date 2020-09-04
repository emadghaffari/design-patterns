package packages

import (
	"fmt"
	"log"
)

func ExampleBar() {
	h1 := Human{
		Name:   "sorena",
		Age:    12,
		Gender: "male",
	}
	bar := OpenBar{}
	// jon is 12 !?
	if err := bar.Welcome(h1); err != nil {
		log.Fatalln("Error in welcome: ", err.Error())
	}
	fmt.Println(bar)
	// Output:
	// {[{sorena 12 male}]}
}
