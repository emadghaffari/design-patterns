package packages

import (
	"fmt"
	"log"
	"testing"
)

func TestMain(t *testing.T) {
	cloner, err := GetTshirClone(WHITE)
	if err != nil {
		log.Fatalln("error: ", err.Error())
	}
	if _, ok := cloner.(*Tshirt); !ok {
		log.Fatalln("error: ", err.Error())
	}
	_ = cloner.GetInfo()
}

func ExampleMain() {
	cloner, err := GetTshirClone(WHITE)
	if err != nil {
		log.Fatalln("error: ", err.Error())
	}
	if _, ok := cloner.(*Tshirt); !ok {
		log.Fatalln("error: ", err.Error())
	}
	item := cloner.GetInfo()
	fmt.Print(item)
	// Output:
	// Tshirt size: 'M', Color: 'white'.

}
