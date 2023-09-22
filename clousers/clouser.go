package clousers

import (
	"fmt"
)

func Double(number int) func(int) int {
	divideBy := 2
	num := 5
	mul := func(num int) int {
		num = num * 2
		fmt.Println("num: in", num)
		return number * divideBy
	}

	fmt.Println("num: out", num)
	return mul
}
