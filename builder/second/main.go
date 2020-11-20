package main

import (
	"fmt"

	"github.com/emadghaffari/design-patterns/builder/second/packages"
)

func main() {
	laptop := packages.LaptopBuilder{}
	builder := packages.Builder{}
	builder.SetBuilder(&laptop)
	laptop.SetCPU("corei 7")
	laptop.SetHDD(1000)
	laptop.SetOS("Windows 7")
	fmt.Println(builder.Build().GetSystem().OS)
}
