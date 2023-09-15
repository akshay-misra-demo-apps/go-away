package main

import (
	"fmt"
	"math/big"
	"reflect"
)

func init() {
	fmt.Println("Init 1")
}

func init() {
	fmt.Println("Init 2")
}

func main() {
	fmt.Println("Hello go")

	y := 10

	z := 10.0

	fmt.Println("Y", reflect.TypeOf(y))
	fmt.Println("Z", reflect.TypeOf(z))

	// Create a new big integer
	bigInt := new(big.Int)

	// Set a value
	bigInt.SetString("5234567890123456789012371678901234567890", 10)

	// Perform arithmetic operations and other operations with bigInt
	fmt.Println(bigInt)
}
