package main

import (
	"fmt"

	"github.com/Emadghaffari/design-patterns/composition/packages"
)

func main() {
	dog := packages.Dog{}
	fmt.Println("Dog:")
	dog.Bark()
	fmt.Println("-------------")

	fmt.Println("CleanerRobotDog:")
	cleanerRobotDog := packages.CleanerRobotDog{}
	cleanerRobotDog.Bark()
	cleanerRobotDog.Clean()
	cleanerRobotDog.Drive()
}
