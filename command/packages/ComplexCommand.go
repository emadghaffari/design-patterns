package packages

import (
	"fmt"
	"log"
)

// ComplexCommand struct
type ComplexCommand struct {
	ID     int
	Status string
}

// Execute method
func (w ComplexCommand) Execute() {
	str, err := w.download()
	if err != nil {
		log.Fatalln("Error in Execute Download Link.")
	}
	str = w.parseHTML(str)
	fmt.Println("executed.", str)
}

// GetID method
func (w ComplexCommand) GetID() int {
	return w.ID
}

// GetStatus method
func (w ComplexCommand) GetStatus() string {
	return w.Status
}

func (w ComplexCommand) download() (string, error) {
	return "download https://google.com", nil
}

func (w ComplexCommand) parseHTML(s string) string {
	return fmt.Sprintf("<h1> %s </h1>", s)
}
