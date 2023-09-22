package routines

import (
	"fmt"
)

// GenerateFibonacci use goroutines, channels, select
func GenerateFibonacci() {
	data := make(chan int)
	done := make(chan bool)
	num := 12
	go func(num int) {
		for i := 0; i <= num; i++ {
			fmt.Println(<-data)
		}
		done <- true
	}(num)

	fibonacci(data, done)
}

func fibonacci(data chan int, done chan bool) {
	a, b := 0, 1

	select {
	case data <- a:
		a, b = b, a+b
	case <-done:
		fmt.Println("series ready")
	}
}
