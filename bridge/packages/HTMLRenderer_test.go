package packages

// Example for Render in HTML Mode
func ExampleHTMLRenderer() {
	r := HTMLRenderer{}
	r.Renderer("Hi from Html tag")
	// Output:
	// html txt: <h1>Hi from Html tag</h1>
}

// Example for interface and HTML Render
func ExampleInterfaceWithHTMLRenderer() {
	render := Render{
		Txt:  "Hello from Me!",
		Rend: &HTMLRenderer{},
	}
	render.Renderer()
	// Output:
	// html txt: <h1>Hello from Me!</h1>
}
