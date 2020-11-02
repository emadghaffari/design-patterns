package main

import (
	"fmt"
	"log"

	"github.com/Emadghaffari/design-patterns/prototype/packages"
)

func main() {
	cloner, err := packages.GetTshirClone(packages.WHITE)
	if err != nil {
		log.Fatalln("error: ", err.Error())
	}
	shirt, ok := cloner.(*packages.Tshirt)
	if !ok {
		log.Fatalln("error: ", err.Error())
	}
	// change M to MS
	shirt.Size = "MS"
	item1 := cloner.GetInfo()

	fmt.Println(item1)
}
