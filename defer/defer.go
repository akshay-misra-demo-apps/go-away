package _defer

import (
	"fmt"
)

func Differ() {
	open()
	defer close()
	flush()
}

func open() {
	fmt.Println("Opening db connection")
}

func close() {
	fmt.Println("Closing db connection")
}

func flush() {
	fmt.Println("Flushing changes to DB")
}
