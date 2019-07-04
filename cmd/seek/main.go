package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, _ := os.Open("test.txt")
	defer file.Close()
	// how many bytes to move
	var offset int64 = 5
	// whence is the point of reference for offset
	// 0 - beginning if file
	// 1 - current position
	// 2 - end of file

	var whence int = 0
	newPosition, err := file.Seek(offset, whence)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Just moved to 5:", newPosition)

	newPosition, err = file.Seek(-2, 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Just moved back two:", newPosition)
}
