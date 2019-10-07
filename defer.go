package main

import "fmt"

func funcdefer() {
	fmt.Println("start func defer")
	n := 10
	defer func() {
		n = n + 10
		fmt.Println(n)
	}()
	defer fmt.Println("defer func 1")
	defer fmt.Println("defer func 2")
	defer fmt.Println("defer func 3")

	fmt.Println(n)
	fmt.Println("finish func defer")
}

func main() {
	fmt.Println("start main")
	funcdefer()
	fmt.Println("finish main")
}
