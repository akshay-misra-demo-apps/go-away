package routines

import (
	"fmt"
	"sync"
)

type singleton struct {
}

var instance *singleton

var lock = &sync.Mutex{}

func New() *singleton {
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		fmt.Println("creating instance!")
		instance = &singleton{}
	} else {
		fmt.Println("instance exists!")
	}

	return instance
}
