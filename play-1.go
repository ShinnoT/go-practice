package main

import "fmt"

func main() {
	fmt.Println("hello world")
	printSomething()
}

// for loops

func printSomething() {
	for i := 0; i < 10; i++ {
		fmt.Printf("this is number %d\n", i)
	}
}
