package packages

import (
	"fmt"
)

// Driver interface
type Driver interface {
	Drive()
}

// Cleaner interface
type Cleaner interface {
	Clean()
}

// Killer interface
type Killer interface {
	Kill()
}

// Barker interface
type Barker interface {
	Bark()
}

// Dog struct
type Dog struct{}

// Bark method
func (d *Dog) Bark() {
	fmt.Println("Wow!")
}

// RobotDog struct
type RobotDog struct{}

// Drive method
func (r *RobotDog) Drive() {
	fmt.Println("Driving...")
}

// Bark method
func (r *RobotDog) Bark() {
	fmt.Println("Wow!")
}

// CleanerRobotDog struct
type CleanerRobotDog struct{}

// Drive method
func (r *CleanerRobotDog) Drive() {
	fmt.Println("Driving...")
}

// Bark method
func (r *CleanerRobotDog) Bark() {
	fmt.Println("Wow!")
}

// Clean method
func (r *CleanerRobotDog) Clean() {
	fmt.Println("Cleaning...")
}

// KillerRobotDog struct
type KillerRobotDog struct{}

// Drive method
func (r *KillerRobotDog) Drive() {
	fmt.Println("Driving...")
}

// Bark method
func (r *KillerRobotDog) Bark() {
	fmt.Println("Wow!")
}

// Kill method
func (r *KillerRobotDog) Kill() {
	fmt.Println("You gonna die!")
}
