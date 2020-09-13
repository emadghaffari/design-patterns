package packages

import (
	"fmt"
	"log"
)

// SimpleCommand struct
type SimpleCommand struct {
	ID     int
	Status string
}

// Execute method
func (w SimpleCommand) Execute() {
	str, err := w.download()
	if err != nil {
		log.Fatalln("Error in Execute Download Link.")
	}
	fmt.Println("executed. ", str)
}

// GetID method
func (w SimpleCommand) GetID() int {
	return w.ID
}

// GetStatus method
func (w SimpleCommand) GetStatus() string {
	return w.Status
}

func (w SimpleCommand) download() (string, error) {
	return "download https://google.com", nil
}
