package packages

// Example for Render in JSON Mode
func ExampleJSONRenderer() {
	r := JSONRenderer{}
	r.Renderer("Hi from JSON tag")
	// Output:
	// {txt: Hi from JSON tag}
}

// Example for interface and JSON Render
func ExampleInterfaceWithJSONRenderer() {
	render := Render{
		Txt:  "Hello from Me!",
		Rend: &JSONRenderer{},
	}
	render.Renderer()
	// Output:
	// {txt: Hello from Me!}
}
