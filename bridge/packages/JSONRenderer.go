package packages

import "fmt"

// JSONRenderer struct
type JSONRenderer struct{}

// Renderer method from Render interface
func (s JSONRenderer) Renderer(txt string) {
	fmt.Printf("{txt: %s}\n", txt)
}
