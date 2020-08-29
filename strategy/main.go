package main

import (
	"flag"
	"log"
	"os"

	"github.gom/emadghaffari/design-patterns/strategy/package/factory"
)

var output = flag.String("output", "console", "the out put can text, image or pdf")

func main() {
	flag.Parse()

	strategy, err := factory.NewPrint(*output)
	if err != nil {
		log.Fatalln(err.Error())
	}

	switch *output {
	case factory.TextStrategy:
		strategy.SetWriter(os.Stdout)
	case factory.ImageStrategy:
		img, err := os.Create("/workspaces/Golang/design-patterns/strategy/image.jpg")
		if err != nil {
			log.Fatalln(err.Error())
		}
		defer img.Close()
		strategy.SetWriter(img)
	case factory.PdfStrategy:
		strategy.SetWriter(os.Stdout)
	default:
		strategy.SetWriter(os.Stdout)
	}

	strategy.Print()
}
