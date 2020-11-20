package packages

import (
	"fmt"
)

// RenderInterface interface
type RenderInterface interface {
	Renderer(Txt string)
}

// Render struct
type Render struct {
	Txt  string
	Rend RenderInterface
}

// Renderer method
func (d Render) Renderer() error {
	if d.Rend == nil {
		return fmt.Errorf("Drawer not initialized")
	}

	d.Rend.Renderer(d.Txt)
	return nil
}
