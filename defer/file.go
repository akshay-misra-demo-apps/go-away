package _defer

import (
	"fmt"
	"log"
	"os"
)

func CreateFile(fileName, content string) {
	fmt.Println("creating a new file with name: ", fileName)

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("error while creating a new file with name: %v", fileName)
	}

	defer file.Close()

	numberOfBytes, err := file.WriteString(content)
	if err != nil {
		log.Fatalf("error while writing to file with name: %v", fileName)
	}

	log.Printf("file created successfully with name: %v and with numberOfBytes %v", file.Name(), numberOfBytes)
}

func ReadFile(filename string) {
	defer Recover()
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Panicf(fmt.Sprintf("file not found with name %v", filename))
	}

	log.Printf("content of file: %v", string(data))
}

func Recover() {
	r := recover()
	if r == nil {
		//log.Println("no panic, all good")
		// do nothing
	} else {
		log.Println("panic recovered")
	}
}
