package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/Emadghaffari/design-patterns/prototype/packages"
)

func TestMain(t *testing.T) {
	cloner, err := packages.GetTshirClone(packages.WHITE)
	if err != nil {
		log.Fatalln("error: ", err.Error())
	}
	if _, ok := cloner.(*packages.Tshirt); !ok {
		log.Fatalln("error: ", err.Error())
	}
	_ = cloner.GetInfo()
}

func ExampleMain() {
	cloner, err := packages.GetTshirClone(packages.WHITE)
	if err != nil {
		log.Fatalln("error: ", err.Error())
	}
	if _, ok := cloner.(*packages.Tshirt); !ok {
		log.Fatalln("error: ", err.Error())
	}
	item := cloner.GetInfo()
	fmt.Print(item)
	// Output:
	// Tshirt size: 'M', Color: 'white'.

}
