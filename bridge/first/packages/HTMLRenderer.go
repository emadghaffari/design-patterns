package packages

import "fmt"

// HTMLRenderer struct
type HTMLRenderer struct{}

// Renderer method from Render interface
func (s HTMLRenderer) Renderer(txt string) {
	fmt.Printf("html txt: <h1>%s</h1> \n", txt)
}
