package main

import (
	"fmt"
	"sync"
)

var (
	singletonInstance *singleton
	once              sync.Once
)

type singleton struct{}

func getSingleton() *singleton {
	once.Do(func() {
		fmt.Println("singletonInstance created.")
		singletonInstance = &singleton{}
	})
	return singletonInstance
}

func main() {

	for i := 0; i < 30; i++ {
		getSingleton()
	}
}
