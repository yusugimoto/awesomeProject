package main

import "fmt"

func funcVariadic(n int, slice ...string) { //可変長引数
	fmt.Print(n, ": ")
	for _, s := range slice {
		fmt.Print(s, ", ")
	}
	fmt.Println()
}

func main() {
	slice := []string{"a", "b"}
	funcVariadic(10, slice...) // スライサーというものを使う
	funcVariadic(10, "a", "b")
}