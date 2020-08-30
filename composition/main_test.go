package main

import (
	"testing"

	"github.com/Emadghaffari/design-patterns/composition/packages"
)

// ExampleDog func for test Dog struct
func ExampleDog(t *testing.T) {
	dog := packages.Dog{}
	dog.Bark()
	// Output:
	// Wow!
}

// ExampleDog func for test Dog struct
func ExampleCleanerRobotDog(t *testing.T) {
	dog := packages.CleanerRobotDog{}
	dog.Bark()
	// Output:
	// Wow!
}
