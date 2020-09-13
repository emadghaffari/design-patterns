package main

import (
	"github.com/Emadghaffari/design-patterns/command/packages"
)

func main() {
	invoker := packages.Invoker{}
	simple := packages.SimpleCommand{
		ID:     2,
		Status: "Active",
	}
	complex := packages.ComplexCommand{
		ID:     931,
		Status: "Deactive",
	}
	invoker.SetOnStart(simple)
	invoker.SetOnFinish(complex)

	invoker.OnStart.Execute()
	invoker.OnFinish.Execute()
}
