package main

import "fmt"

// Arrays
// The type [n]T is an array of n values of type T.
// The expression : 
//    var a [10]int
// declares a variable a as an array of ten integers.

func main() {
	
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	fmt.Println("-----> int:")
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	fmt.Println("-----> str:")
	str := [6]string{"php", "java", "javascript", "html"}
	fmt.Println(str)
}
