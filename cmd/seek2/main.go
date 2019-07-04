package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = file.Write([]byte("hello"))
	if err != nil {
		log.Fatal(err)
	}

	offset, err := file.Seek(0, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(offset)

	_, err = file.Write([]byte("world"))
	if err != nil {
		log.Fatal(err)
	}
	offset, err = file.Seek(0, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(offset)
}
