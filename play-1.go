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

// for loops in go
i := 1
for i <= 10 {
	fmt.Println(i)
	i += 1
}

for i := 1; i <= 10; i++ {
	fmt.Println(i)
}

// simple if else if else
if i % 2 == 0 {
	// even
} else if != 0 {
	// odd
} else {
	// NaN
}

// arrays in go
// fixed length with only one type elements
var somearray [5]int
somearray[4] = 100
fmt.Println(somearray)
// prints [0 0 0 0 100]

func average() float64 {
	var z [5]int
	z[0] = 20
	z[3] = 32
	var total float64 = 0
	for i := 0; i < len(z); i++ {
		total += z[i]
	}
	// will result in error since total and len(z) is different types so...
	// return total / len(z)
	return total / float64(len(z))
}

// another way of using for loop, similar to a forEach in Node.js
var z [5]int
z[0] = 20
z[3] = 32
// below will give error warning because we decalare the variable index but we dont use it
// for index, value := range z {
// 	fmt.Println(value)
// 	// will print each value
// }

// so....
for _, value := range z {
	fmt.Println(value)
}
// THE UNDERSCORE IN THIS CASE WILL
// ALLOW YOU TO HAVE AN UNDECLARED VARIABLE IN PLACES YOU NEED TO PUT IT

// go short syntax for creating arrays
somearray := [5]float64{98, 88, 33, 55, 63}
// you can put it on different lines as such:
secondarray := [3]float64{
	33,
	21,
	15,
}
//////////////!!!!!!!!! last railing comma needed when doing multi line array declaration

// since array is fixed length, removing elements becomes mendou cuz u gotta change size of array too
// SO WE USE SLICES
// which is a type built on an array

var someSplice []float64

otherSlice := make([]float64, 5)
// this is a slice with an associated array of length 5 (it cant be longer but can always be shorter in length)

// the makes function also takes in a 3rd parameter
evenMoreSlice := make([]float64, 5, 10)

// another way to create slices is to use [low : high] expression
arr := [5]float64{1, 2, 3, 4, 5}
aS := arr[0:5]
// aS is a slice created from the arr array

bS := arr[0:]
cS := arr[:5]
dS := arr[:]
lS := arr[2:4]

// 2 built in functions to help with slices

// 1- append
slice1 := []int{1,2,3}
slice2 := append(slice1, 4, 5)
// slice1 == [1, 2, 3]
// slice2 == [1, 2, 3, 4, 5]

// 2- copy
slice3 := []int{1, 2, 3}
slice4 := make([]int, 2)
copy(slice4, slice3)
// slice3 == [1, 2, 3]
// slice4 == [1, 2]
// it copied slice3 elements into slice4 however since slice4 only has length 2
// only first two elements were coppied

// MAPS == JS objects ---> key value pair
// maps have no fixed length
// maps are not sequential meaning they have no order


var map1 map[string]int
// map1 is a map with string keys and int values

// settiing key value pair
map1["first key"] = 22
// above wont work because the map was not INITIALIZED
// all maps have to be initialized
map2 := make(map[string]int)
map2["some key"] = 99

// built in map functions

// delete
delete(map2, "some key")

name, ok := map2["some key"]
// name returns the value of the specified key
// ok returns if there was a value associated to the key (boolean)

// typically used like this:
if name, ok := map2["some key"]; ok {
	fmt.Println(name + " was found!")
}

// short way of creating maps
randomMap := map[string]string{
	"name": "shinno",
	"nickname": "snoop dogg",
}
// AGAIN WITH THE LAST TRAILING COMMA

// map of maps

peopleMap := map[string]map[string]int{
	"shinno": map[string]int{
		"age": 25,
		"height": 17999999,
	},
	"doug": map[string]int{
		"age": 25,
		"height": 17999999,
	},
}