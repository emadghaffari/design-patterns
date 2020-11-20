package main

import (
	"github.com/Emadghaffari/design-patterns/bridge/packages"
)

func main() {
	render := packages.Render{
		Txt:  "Hello from Me!",
		Rend: &packages.HTMLRenderer{},
	}
	render.Renderer()

	render2 := packages.Render{
		Txt:  "this is me!",
		Rend: &packages.JSONRenderer{},
	}
	render2.Renderer()

}
