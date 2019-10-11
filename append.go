package main

import "fmt"

func main() {
	slice := []int {1,2,3}
	fmt.Println("slice: ", slice)
	slice2 := append(slice, 4) // 要素の最後1を追加する
	fmt.Println("slice2: ", slice2)
	slice3 := append(slice, slice...) // 配列同士の結合 第2引数に「スライス型...」 を指定する「...」も含むので注意
	fmt.Println("slice3: ", slice3)
}
