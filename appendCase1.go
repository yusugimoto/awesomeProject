package main

import "fmt"

func main() {
	fmt.Println("空き容量がある場合")
	slice := make([]int, 2, 3) // スライスの長さ3
	slice2 := append(slice, 1) // 空きがあるのでスライスの参照渡しが行われる

	fmt.Println("slice: ", slice)
	fmt.Println("slice2: ", slice2)
	slice[1] = 10 // slice, slice1両方の値が変わる

	fmt.Println("slice: ", slice)
	fmt.Println("slice2: ", slice2)


	fmt.Println("空き容量が足りない場合")
	slice3 := make([]int, 2, 2) // スライスの長さ2
	slice4 := append(slice3, 1) // 空きがないのでコピーされる

	fmt.Println("slice3: ", slice3)
	fmt.Println("slice4: ", slice4)
	slice3[1] = 10 // slice4はslice3をコピーして作成されるのでslice4には影響を与えない

	fmt.Println("slice3: ", slice3)
	fmt.Println("slice4: ", slice4)
}
