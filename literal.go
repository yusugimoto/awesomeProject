package main

import "fmt"

func main() {
	f := func(i, j int) int {
		return i + j
	}
	n := f(1 ,1)
	fmt.Println(n)
}
