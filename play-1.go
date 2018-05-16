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


// functions

func averageSecond(someSlice []float64) float64 {
	var total float64 = 0
	for i := 0; i < len(someSlice); i++ {
		total += someSlice[i]
	}
	// will result in error since total and len(someSlice) is different types so...
	// return total / len(someSlice)
	return total / float64(len(someSlice))
}

// in go we can name return types
func f2() (r int) {
	r =1
	return
}

// WE CAN RETURN MULTIPLE VALUES IN GO!!!!
func foo() (int, int) {
	return 5, 6
}

// this way you can get two values when you call the function later in main like this
// func main() {
	firsVal, secondVal := foo()
// }

// in go return multiple values is used in return the actual value and error or boolean to indicate success

// variadic functions, having one parameter but passing multiple arguments

func add(args ...int) int {
	total := 0
	for _, value := range args {
		total += value
	}
	return total
}
// so when you call add() you can do
add(1, 2, 3) // passing in more than one argument
// we can even pass slice so long that we trail it with elipses

randomSl := []int{1, 2, 3}
add(randomSl...)

// functions as variables like arrow functions!!!!!!!!!! (also function within functions)
func main() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(2, 5))
}

// HIGHER ORDER FUNCTIONNSSSSSS func that returns a func
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() uint {
		ret = i
		i += 2
		return ret
	}
}

// RECURSIONNNNNNNNNNNN FUNCTION CALLING ITSELF
// factorial
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x - 1)
}

// special statements in GO for functions

// defer - schedules a function call to be run after the fucntion compiles
func first() {
	fmt.Println("first")
}

func second() {
	fmt.Println("second")
}

func main() {
	defer second()
	first()
}

// will print first then second even though second is called first because it is being deferred meaning it is being told to wait until function compiles

// mostly used in situatioins where resources need to be freed in some way
f, _ := os.Open(filename)
defer f.Close()

// PANIC AND RECOVER - recovering from a runtime error
// *** small not ::: u can directly call (only nameless funcs???????) function while defining it by adding () at the end or (with args here)
func() {
	// do something
}() // calls it here


// say we have a func that calls the builtin panic func that will throw some runtime error
// if it is at the top of the call stack, once its been compiled it will throw error and so rest of code wont run
// how can we stop this from happening so that even if first on callstack has error, we still recover and run rest of code?
// USING DEFER AND RECOVER

func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}

// POINTERS

// consider the examples below
func zero(x int) {
	x = 0
}

func main() {
	x := 5
	zero(x)
	fmt.Println()
}

// main function will still print 5 here because it copies the variable argument within the function and so it doesnt mutate the original argument
// we can modify original argument if we want to by using special data type called pointers
func zero(xPtr *int) {
	*xPtr = 0
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x)
}

// * will point to a memory location specified and access or write something in that memory space
// & is to find the memory space of a variable (it returns a pointer to int)

// new way of doing it is using new() built in function
func one(xPtr *int) {
	*xPtr = 1
}

func main() {
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr) // x is 1
}

// new() will take type as arg and allocates enough memory to fit a value of that type, and returns pointers to it
// in GO new() and & is the same unlike other languages
// it is a garbage collecting language so memory is automatically cleaned up when not in use


// ---------------------------------
// STRUCTS and INTERFACES

// structs similar to classes
type Circle struct {
	x float64
	y float64
	r float64
}
// short hand:
type Circle struct {
	x, y, r float64
}

// creating instances of new Circle type:
var c Circle
// a more uncommon way below
c := new(Circle)

// how to set initial values and not just empty instances?
c := Circle{x: 0, y: 0, r: 5}
// if we know the order the fields were defined we can do
c := Circle{0, 0, 5}
// if you want a pointer to the struct use &
c := &Circle{0, 0, 5}

// we can access fields of instances using . notation
fmt.Println(c.r)
c.x = 10
// etc.

// if we want to modify fields of instances of Circle, obvie gotta use pointer
func circleArea(c *Circle) float64 {
	return math.Pi * c.r*c.r
}

func main() {
	c := Circle{0,0,5}
	fmt.Println(circleArea(&c))
}

// Methods - special type of functions called methods
// instance methods for the structs you made!!!!!!

// before the name of the function, we specify which struct to attach the function to
func (circleInstance *Circle) area() float64 {
	return math.Pi * circleInstance.r*circleInstance.r
}
// now we can call it using . notation on the instance of the struct
fmt.Println(c.area())


// embedded types
type Person struct {
	Name string
}
func (p *Person) Talk() {
	fmt.Println("Hi my name is ", p.Name)
}
// type Employee struct {
// 	Person Person
// 	Position string
// }
// but this says employee HAS a person not IS a person
// so instead we can do
type Employee struct {
	Person
	Position string
}
randomEmployee := Person{Position: "software engineer"} //?????
randomEmployee.Person.Talk()
randomEmployee.Talk()

// Interface
// area() is common to all shapes fot ex. so we make parent method interface to apply to all shapes instead of defining it for each
type Shape interface {
	area() float64
}
func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}
// if we have c Circle and r Radius
fmt.Println(totalArea(&c, &r))
// all that the totalArea knows is that there is an area() method for each Shape
// so it can only access area() of each shape struct such as Circle or Rectangle
// it CANNOT access anything else like the radius for example

// interfaces can also be used as fields
type MultiShape struct {
	shapes []Shape
}

multiShapeOne := MultiShape{
	shapes: []Shape{
		Circle{0,0,5},
		Rectangle{0,0,10,10},
	},
}

// interfaces become useful as the program develops
// allows us to hide incidental details of implementation
// in our example, as long as the area() method is defined in all shape structs like Circle or Rectangle
// we are free to do whatever with other fields without having to modify the interface method

// ----------------------------
// PACKAGES
// packages include a variety of reusable functions
	// reduces change of overlappnig names, in turn keeping function names short
	// organizes code so easy to find what u want to use
	// speeds up compiler by only requiring recompilation of smaller chunks of a program


// core packages:

import (
	"fmt"
	"strings"
	"io"
	"bytes"
	"os"
	"io/ioutil"
)

func main() {
	// string manipulation package "string"
	fmt.Println(strings.Contains("test", "es")) // true
	fmt.Println(strings.Count("test", "t")) // 2
	fmt.Println(strings.HasPrefix("test", "te")) // true
	fmt.Println(strings.HasSuffix("test", "st")) // true
	fmt.Println(strings.Index("test", "e")) // 1
	fmt.Println(strings.Join([]string{"a", "b"}, "-")) // "a-b"
	fmt.Println(strings.Repeat("a", 5)) // "aaaaa"
	fmt.Println(strings.Replace("aaaa", "a", "b", 2)) // "bbaa" (put -1 to have it repeat for every accurance)
	fmt.Println(strings.Split("a-b-c", "-")) // []string{"a", "b", "c"}
	fmt.Println(strings.ToLower("TEST")) // "test"
	fmt.Println(strings.ToUpper("test")) // "TEST"
	// work with strings as binary data. string to slice of binary data and vise versa
	arr:= []byte("test")
	str := string([]byte{'t', 'e', 's', 't'})

	// input output package "io"
	// consists of few functions but mostly interfaces used in other packages
	// most notably the Reader and Writer interfaces

	// to read or write to []byte or string, use Buffer struct in bytes package
	var buf bytes.Buffer
	buf.Write([]byte("test"))
	// you can convert buf into []byte by calling buf.Bytes()
	// if you only need to read from string use strings.NewReader --> more efficient than Buffer

}

func filesAndFolder() {
	// files and folders package "os"
	// here is code to open contents of file and print on terminal
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("sorry there was an error")
		return
	}
	defer file.Close() // we use this line to ensure that the file is complete only right after function completes

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		fmt.Println("error again sorry")
		return
	}
	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		fmt.Println("lol another error")
		return
	}
	str := string(bs)
	fmt.Println(str)
}

// since reading files is common, theres another shorter way to do the above
func shorterReadFile() {
	bs, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("error dangerrrr")
		return
	}
	str := string(bs)
	fmt.Println(str)
}

// creating a file
func writeFile() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("omg error")
		return
	}
	defer file.Close()

	file.WriteString("test")
}

// open contents of directory
func openDirectory() {
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()

	// Readdir takes one argument, a int that limits the number of files returned
	// by passing in -1 we return ALL the files
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}

// sometimes we want to recursively crawl (walk) a folder
// read a folder, its files, subfolders, then subfolders of those folders, etc. etc.
// to do this, there is a path/filepath package

import {
	"fmt"
	"os"
	"path/filepath"
}

func readAllFoldersAndSubfolders() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
// the function passed to filepath.Walk() will be called for every file and folder in the root folder
// filepath.SkipDir ?????


// Error type and creating custom errors
import "errors"

func main() {
	err := errors.New("custom error message")
}

// CONTAINERS AND SORT

// container/list package implements a doubly linked list
// each node of the list contains a value and a pointer to the next node
// since doubly linked list, each node will also have pointers to the previous node
// we can create a doubly linked list as such:

import (
	"fmt"
	"container/list"
)

func main() {
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	for e := x.Front(); e != nil; e=e.Next() {
		fmt.Println(e.Value.(int))
	}
}
// we print every value of each node by iterating starting from the first node in the linked list


// SORT!!!!!!
// sort package contains functions for sorting arbitrary data
// predifined sorts for slices

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age int
}

type ByName []Person

func (ps ByName) Len() int {
	return len(ps)
}

func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}

func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func main() {
	kids := []Person{
		{"Jill",9},
		{"Jack",10},
	}
	sort.Sort(ByName(kids))
	fmt.Println(kids)
}
// Len, Less, and Swap functions are necessary for sort algorithms


// HASHES AND CRYPTOGRAPHY
// takes set of data and reduces it to a smaller fixed size
// two types: cryptographic and non-cryptographic

// non-cryptographic found under hash package
// adler32, crc32, crc64, and fnv
import (
	"fmt"
	"hash/crc32"
)

func main() {
	// create hasher
	h := crc32.NewIEEE()
	// write our data to it
	h.Write([]byte("test"))
	// calculate the crc32 checksum
	v := h.Sum32()
	fmt.Println(v)
}

// cyptographic hashing is the same but made almost irrevirsible
// so that one cannot reproduce the hash
import (
	"fmt"
	"crypto/sha1"
)

func main() {
	h := sha1.New()
	h.Write([]byte("test"))
	bs := h.Sum([]byte{})
	fmt.Println(bs)
}
// same as non-cryptgraphic hash except crc32 computes 32-bit hash, sha1 computes 160-bit hash
// no native type to represent a 160-bit number, so we use slice of 20bytes instead



// SERVERS
// TCP - primary protocol for communication over the internet

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server() {
	// listen on a port
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		dmt.Println(err)
		return
	}
	for {
		// accept a connection
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// handle the connection (function defined below)
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	// receive the message
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received", msg)
	}
	c.Close()
}

func client() {
	// connect to the server
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := "Hello, World"
	fmt.Println("Sending", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}

	c.Close()
}

func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}