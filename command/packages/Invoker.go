package packages

// Invoker struct
type Invoker struct {
	OnStart  Command
	OnFinish Command
}

// SetOnStart method
func (i *Invoker) SetOnStart(c Command) {
	i.OnStart = c
}

// SetOnFinish method
func (i *Invoker) SetOnFinish(c Command) {
	i.OnFinish = c
}
