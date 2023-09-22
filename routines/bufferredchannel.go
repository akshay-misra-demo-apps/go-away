package routines

import (
	"fmt"
)

// Bufferred normal channel is one send and one receive
// until we read from the channel it blocks the control
// but with buffer until the buffer limit reached, it will not block the control
func Bufferred() {
	// buffer channel with a cap of 5
	// until channel accepts 5 int, the control will not be blocked
	ch := make(chan int, 10)

	for i := 0; i < 10; i++ {
		ch <- i
	}

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}
