package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println(os.Args[1])

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("Error: ", err)
	}

	io.Copy(os.Stdout, f)

}
