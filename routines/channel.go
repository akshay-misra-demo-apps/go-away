package routines

import (
	"fmt"
	"time"
)

// chan int: bi directional channel
// chan <- int sender channel
// <- chan int receive only channel

func ExecuteChannel() {
	fmt.Println("Execute channel starts")

	buffer := make(chan string)
	defer close(buffer)
	done := make(chan bool)
	defer close(done)
	go writer(buffer, done)
	//close(buffer)
	reader(buffer, done)
}

func writer(buffer chan string, done chan bool) {
	fmt.Println("writing starts!")

	buffer <- "hello"
	time.Sleep(time.Second)
	buffer <- "how"
	time.Sleep(time.Second)
	buffer <- "are"
	time.Sleep(time.Second)
	buffer <- "you"
	time.Sleep(time.Second)

	done <- true

	fmt.Println("writing ends!")
}

func reader(buffer chan string, done chan bool) {

	//for {
	//	if msg, ok := <-buffer; ok {
	//		fmt.Println("channel is open: got data: ", msg)
	//	} else {
	//		fmt.Println("channel is closed: got data: ", msg)
	//		return
	//	}
	//}

	for {
		select {
		case msg := <-buffer:
			fmt.Println("Received:", msg)
		case <-done:
			fmt.Println("Received done signal. Exiting!")
			return
		}
	}
}

// Generator       receiver channel
func Generator() <-chan int {
	buffer := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {

			fmt.Println("generating ", i)
			buffer <- i
		}
		close(buffer)
	}()

	return buffer
}

func Receiver(buffer <-chan int) {
	for i := range buffer {
		fmt.Println("receiving: ", i)
	}
}
