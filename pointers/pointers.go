package pointers

import (
	"fmt"
	"reflect"
)

func Pointers() {
	var i = 10
	var p = &i
	fmt.Println("value of i: ", i)
	fmt.Println("address of i: ", p)
	fmt.Println("value stored in pointer to i: ", *p)
	fmt.Println("data type of pointer: ", reflect.TypeOf(p))
	fmt.Println("data type of  value stored in pointer to i: ", reflect.TypeOf(*p))
}
