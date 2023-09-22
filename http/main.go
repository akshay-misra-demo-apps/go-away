package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	fmt.Println("server started at port 4000")
	if err := http.ListenAndServe(":4000", nil); err != nil {
		log.Fatalf("error while starting the http server: %v", err.Error())
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Welcome to advance golang training!</>")
}
