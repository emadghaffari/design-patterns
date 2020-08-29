package singleton

// Singleton interface
type Singleton interface {
	AddOne() int
}

// singleton struct
type singleton struct {
	counter int
}

var instance *singleton

// GetInstance func
func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

// AddOne method
func (co *singleton) AddOne() int {
	co.counter++
	return co.counter
}
