package packages

import "fmt"

// Adapter struct
type Adapter struct {
	Adapter *Windows
}

// InsertIntoLightningPort meth
func (a *Adapter) InsertIntoLightningPort()  {
	fmt.Println("adapter for windows")
    a.Adapter.InsertIntoUSBPort()
}