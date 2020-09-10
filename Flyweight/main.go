package main

import (
	"log"

	"github.com/Emadghaffari/design-patterns/flyweight/packages"
)

func main() {
	factory := packages.NewTeamFactory()
	Team1 := factory.GetTeam(packages.Team_A)
	if Team1 == nil {
		log.Fatalln("Error in Team1.")
	}
	Team2 := factory.GetTeam(packages.Team_A)
	if Team2 == nil {
		log.Fatalln("Error in Team2.")
	}

	if Team1 != Team2 {
		log.Fatalln("Error in Team1 != Team2")
	}

	if factory.GetNumberOfObject() != 1 {
		log.Fatalln("Error in get number of object")

	}
}
