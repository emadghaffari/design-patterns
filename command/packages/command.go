package packages

// Command interface
type Command interface {
	Execute()
	GetID() int
	GetStatus() string
}
