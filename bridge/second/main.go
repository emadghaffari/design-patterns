package main

import "github.com/emadghaffari/design-patterns/bridge/second/packages"

func main() {
	hp := packages.Hp{}
	epson := packages.Epson{}

	windows := packages.Windows{}
	mac := packages.Mac{}

	windows.SetPrinter(&hp)
	windows.Print()

	windows.SetPrinter(&epson)
	windows.Print()

	mac.SetPrinter(&epson)
	mac.Print()

	mac.SetPrinter(&hp)
	mac.Print()
}
