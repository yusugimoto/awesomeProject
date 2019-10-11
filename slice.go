package main

import "fmt"

func main() {
	slice := []int {1,2,3}
	slice2 := slice //参照渡し
	slice[1] = 10 //参照渡しなのでslice2[1]も値が変わる

	fmt.Println("slice: ", slice)
	fmt.Println("slice2: ", slice2)
}
