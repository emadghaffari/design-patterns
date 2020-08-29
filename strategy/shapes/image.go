package shapes

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"

	"github.gom/emadghaffari/design-patterns/strategy/package/strategy"
)

// ImageSquare struct
type ImageSquare struct {
	strategy.PrintOutPut
}

// Print method
func (i ImageSquare) Print() error {
	with := 1200
	height := 900

	bgColor := image.Uniform{color.RGBA{R: 70, G: 70, B: 70, A: 0}}
	origin := image.Point{0, 0}
	quality := jpeg.Options{Quality: 75}

	bgImage := image.NewRGBA(image.Rectangle{
		Min: origin,
		Max: image.Point{X: with, Y: height},
	})
	draw.Draw(bgImage, bgImage.Bounds(), &bgColor, origin, draw.Src)

	squareWidth := 200
	squareHeight := 200
	squareColor := image.Uniform{color.RGBA{0, 255, 0, 0}}
	square := image.Rect(0, 0, squareWidth, squareHeight)
	square = square.Add(image.Point{
		X: (with / 2) - (squareWidth / 2),
		Y: (height / 2) - (squareHeight / 2),
	})
	squareIMG := image.NewRGBA(square)

	draw.Draw(bgImage, squareIMG.Bounds(), &squareColor, origin, draw.Src)

	squareWidth = 100
	squareHeight = 100
	squareColor = image.Uniform{color.RGBA{255, 0, 0, 0}}
	square = image.Rect(0, 0, squareWidth, squareHeight)
	square = square.Add(image.Point{
		X: (with / 2) - (squareWidth / 2),
		Y: (height / 2) - (squareHeight / 2),
	})
	squareIMG = image.NewRGBA(square)

	draw.Draw(bgImage, squareIMG.Bounds(), &squareColor, origin, draw.Src)

	if i.Writer == nil {
		return fmt.Errorf("eror in i.Writer")
	}

	if err := jpeg.Encode(i.Writer, bgImage, &quality); err != nil {
		return err
	}

	return nil
}
