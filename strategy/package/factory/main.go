package factory

import (
	"fmt"
	"os"

	"github.gom/emadghaffari/design-patterns/strategy/package/strategy"
	"github.gom/emadghaffari/design-patterns/strategy/shapes"
)

// textStrategy const
// imageStrategy const
const (
	TextStrategy  = "text"
	ImageStrategy = "image"
	PdfStrategy   = "pdf"
)

// NewPrint method
func NewPrint(s string) (strategy.PrintStrategy, error) {

	switch s {
	case TextStrategy:
		return &shapes.TextSquare{
			PrintOutPut: strategy.PrintOutPut{
				LogWriter: os.Stdout,
			},
		}, nil

	case ImageStrategy:
		return &shapes.ImageSquare{
			PrintOutPut: strategy.PrintOutPut{
				LogWriter: os.Stdout,
			},
		}, nil

	case PdfStrategy:
		return &shapes.PdfSquare{
			PrintOutPut: strategy.PrintOutPut{
				LogWriter: os.Stdout,
			},
		}, nil
	default:
		return nil, fmt.Errorf("error in factory")
	}
}
