package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("Hello, World!")
	log.Println("Hello, World!")

	fmt.Printf("Unix timestamp:  %d \n", time.Now().Unix())
}
