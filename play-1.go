package main

import "fmt"

func main() {
	fmt.Println("hello world")
	// printSomething()
}

// for loops

// func printSomething() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("this is number %d\n", i)
// 	}
// }

// go built in data types
// ** numbers like 32, 64, etc. represent bits

// integer types: uint8, uint16, uint32, uint64, int8, int16, int32, int64 (int type mostly used)
// float types: float32 float64
// complex number types: complex128, complex64 (generally float64 most commonly used)

// declaring variables
// go knows basic types so it automatically infers them without manually declaring type

var x string = "hello world"
var y = "hello world"
z := "hello world"

var w string
w = "hi there"

// constant variables unmutable after they have been made
const a string = "hello world"

// shortcut to defin multiple variables at once
var (
	c = 5
	b = 20
)

const (
	d = 30
	something = "hi"
)

