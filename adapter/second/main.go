package main

import "github.com/emadghaffari/design-patterns/adapter/second/packages"


func main()  {

	mac := &packages.Mac{}
	client := packages.Client{}
	client.InsertLightningConnectorIntoComputer(mac)

	windows := &packages.Windows{}
	windowsADP := &packages.Adapter{
		Adapter: windows,
	}
	client.InsertLightningConnectorIntoComputer(windowsADP)
}