package packages

import "fmt"

// Windows struct
type Windows struct{}

// InsertIntoUSBPort meth
func (w *Windows) InsertIntoUSBPort() {
    fmt.Println("USB connector is plugged into windows machine.")
}