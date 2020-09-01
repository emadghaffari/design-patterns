package main

import (
	"fmt"

	"github.com/Emadghaffari/design-patterns/adapter/packages/new"

	"github.com/Emadghaffari/design-patterns/adapter/packages/old"
)

func main() {
	oldPrinter := old.MyPrinter{}
	oldMsg := oldPrinter.Print("The old message")
	fmt.Println(oldMsg)

	newPrinter := new.MyPrinter{"new Message from old Printer", oldPrinter}
	newMsg := newPrinter.Print()
	fmt.Println(newMsg)
}
