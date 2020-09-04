package packages

import "fmt"

// ExampleWelcome func
func ExampleWelcome() {
	h1 := Human{
		Name:   "jon",
		Age:    15,
		Gender: "male",
	}
	proxyBar := BarProxy{}
	if err := proxyBar.Welcome(h1); err != nil {
		fmt.Println("Error in welcome: ", err.Error())
	}
	// Output:
	// Error in welcome:  Males under 18 not allowed
}

// ExampleSuccessWelcome func
func ExampleSuccessWelcome() {
	h1 := Human{
		Name:   "jon",
		Age:    55,
		Gender: "male",
	}
	proxyBar := BarProxy{}
	if err := proxyBar.Welcome(h1); err != nil {
		fmt.Println("Error in welcome: ", err.Error())
	}

	// Output:
	//
}
