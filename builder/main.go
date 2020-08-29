package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.gom/emadghaffari/design-patterns/builder/packages"
)

func main() {
	manufacturer := packages.Manufacturer{}
	homePCBuilder := &packages.HomeEditionPCBuilder{}

	manufacturer.SetBuilder(homePCBuilder)
	manufacturer.ConstructPC()

	jhome, err := json.Marshal(homePCBuilder.GetPC())
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(string(jhome))
}
