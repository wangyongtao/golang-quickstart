package main

import (
	"fmt"
	"time"
	"log"
)

func main() {
    fmt.Println("Hello, World!");
    log.Println("Hello, World!");

    fmt.Printf("Unix timestamp:  %d \n", time.Now().Unix())
}