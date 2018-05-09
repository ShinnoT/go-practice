package main

import "fmt"

func main() {
	// arr := [5]float64{1, 2, 3, 4, 5}
	// x := arr[0:5]
	// fmt.Println(x)
	map1 := make(map[string]int)
	map1["some key"] = 5
	map1["another key"] = 2
	name, ok := map1["some key"]
	fmt.Println(name, ok)
}