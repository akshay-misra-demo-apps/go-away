package routines

import (
	"fmt"
	"sync"
)

func Execute() {
	fmt.Println("Execute starts")

	// new keyword reserves memory in heap, wg variable will be in stack
	// scope is beyond func, until garbage collector clears it.
	//wg := new(sync.WaitGroup)

	// this is more idiomatic way to create a struct as in this case
	// both struct and variable will be in stack

	// prefer this way over 'new' func to have stack scope, it clears memory as func exits.
	var wg = sync.WaitGroup{}
	// There is a difference in scope of wg in both above case.
	wg.Add(4)
	go hello(&wg)
	go how(&wg)
	go are(&wg)
	go you(&wg)

	wg.Wait()
	fmt.Println("Execute ends")
}

func hello(wg *sync.WaitGroup) {
	defer wg.Done()
	//wg.Add(1)
	fmt.Println("Hello")
}

func how(wg *sync.WaitGroup) {
	defer wg.Done()
	//wg.Add(1)
	fmt.Println("how")
}

func are(wg *sync.WaitGroup) {
	defer wg.Done()
	//wg.Add(1)
	fmt.Println("are")
}

func you(wg *sync.WaitGroup) {
	defer wg.Done()
	//wg.Add(1)
	fmt.Println("you!")
}
