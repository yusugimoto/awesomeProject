package main

import "fmt"

func main() {
	array := [3]int {1,2,3}
	array2 := array
	array[1] = 10

	fmt.Println("array: ", array)
	fmt.Println("array2: ", array2)
}
