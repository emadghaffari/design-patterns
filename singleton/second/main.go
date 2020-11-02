package main

import (
	"fmt"
	"sync"
)

var (
	singletonInstance *singleton
	lock = &sync.Mutex{}
)
type singleton struct{}

func getSingleton() *singleton {
	if singletonInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singletonInstance == nil{
			fmt.Println("singletonInstance created.")
			singletonInstance = &singleton{}
		} else {
			fmt.Println("singletonInstance is already created.")

		}
	} else {
		fmt.Println("singletonInstance is already created.")

	}
	return singletonInstance
}

func main()  {
	
	for i := 0; i < 30; i++ {
		getSingleton()
	}
}