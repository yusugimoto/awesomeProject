package main

import "fmt"

func main() {
	n := 10
	f := func() {
		n = n + 5
	}
	f()
	fmt.Println(n)
}
