package main

import (
	"flag"
	"log"
	"os"
	"testing"

	"github.gom/emadghaffari/design-patterns/strategy/package/factory"
)

// TestMain func
// go test -run '' -output=text
// go test -run '' -output=image
// go test -run '' -output=file
func TestMain(m *testing.M) {
	flag.Parse()
	strategy, err := factory.NewPrint(*output)
	if err != nil {
		log.Fatalln(err.Error())
	}

	switch *output {
	case factory.TextStrategy:
		strategy.SetWriter(os.Stdout)
		strategy.Print()
	case factory.ImageStrategy:
		img, err := os.Create(imagePath)
		if err != nil {
			log.Fatalln(err.Error())
		}
		defer img.Close()
		strategy.SetWriter(img)
		strategy.Print()

		if _, err := os.Stat(imagePath); os.IsNotExist(err) {
			log.Fatalln("the file not Exist: ", err.Error())
		}
	case factory.PdfStrategy:
		strategy.SetWriter(os.Stdout)
		strategy.Print()

		if _, err := os.Stat(pdfPath); os.IsNotExist(err) {
			log.Fatalln("the file not Exist: ", err.Error())
		}
	default:
		strategy.SetWriter(os.Stdout)
		strategy.Print()

	}

}

// ExampleMain
func ExampleMain() {
	flag.Parse()
	strategy, err := factory.NewPrint(*output)
	if err != nil {
		log.Fatalln(err.Error())
	}
	strategy.SetWriter(os.Stdout)
	strategy.Print()
	// Output:
	// circle
}

func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		flag.Parse()
		strategy, err := factory.NewPrint(*output)
		if err != nil {
			log.Fatalln(err.Error())
		}
		strategy.SetWriter(os.Stdout)
		strategy.Print()
	}
}
