package main

import "fmt"

func funcVariadic(n int, slice ...string) string { //可変長引数
	fmt.Print(n, ": ")
	for _, s := range slice {
		fmt.Print(s, ", ")
	}
	fmt.Println()
	return "true"
}

func main() {
	slice := []string{"a", "b"}
	funcVariadic(10, slice...) // スライサーというものを使う
	fmt.Println(funcVariadic(10, "a", "b"))
	fmt.Println(funcResult())
}

func funcResult() (result int) {
	result = 100
	return 20
}