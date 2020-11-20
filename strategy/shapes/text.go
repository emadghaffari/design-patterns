package shapes

import "github.gom/emadghaffari/design-patterns/strategy/package/strategy"

// TextSquare struct
type TextSquare struct {
	strategy.PrintOutPut
}

// Print method
func (t *TextSquare) Print() error {
	t.Writer.Write([]byte("circle\n"))
	return nil
}
