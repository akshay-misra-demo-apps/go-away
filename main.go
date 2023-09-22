package main

import (
	"github.com/akshay-misra-demo-apps/go-away/workerpool"
)

func init() {
	//fmt.Println("Init 1")
}

func init() {
	//fmt.Println("Init 2")
}

func main() {
	//fmt.Println("Hello go")

	//y := 10
	//z := 10.0
	//fmt.Println("Y", reflect.TypeOf(y))
	//fmt.Println("Z", reflect.TypeOf(z))
	//
	//// Create a new big integer
	//bigInt := new(big.Int)
	//// Set a value
	//bigInt.SetString("5234567890123456789012371678901234567890", 10)
	//// Perform arithmetic operations and other operations with bigInt
	//fmt.Println(bigInt)
	//
	//mul := clousers.Double(10)
	//fmt.Println(mul(5))
	//fmt.Println(mul(5))
	//fmt.Println(mul(5))
	//
	//_defer.Differ()
	//
	//_defer.CreateFile("test.txt", "Hello how are you!")
	//
	//_defer.ReadFile("test.txt")
	//
	//fmt.Println("continuing main flow")
	//
	//pointers.Pointers()

	// doubly linked list of type int
	//list := collections.NewDoublyLinkedList[int]()
	//
	//list.Add(1)
	//list.Add(4)
	//list.Add(7)
	//list.Add(0)
	//
	//list.String()
	//
	//list.RemoveByIndex(2)
	//
	//list.String()
	//
	//// doubly linked list of type string
	//list1 := collections.NewDoublyLinkedList[string]()
	//list1.Add("Kiran")
	//list1.Add("Matti")
	//list1.Add("Akshay")
	//list1.Add("Kimmo")
	//
	//list1.String()
	//
	//list1.RemoveByIndex(2)
	//
	//list1.String()

	//routines.Execute()
	//fmt.Println("main ends")

	//routines.ExecuteChannel()

	//buffer2 := routines.Generator()
	//routines.Receiver(buffer2)

	//routines.Bufferred()

	//for i := 0; i < 10; i++ {
	//	go routines.New()
	//}

	//routines.GenerateFibonacci()

	workerpool.ProduceJobs()
}
