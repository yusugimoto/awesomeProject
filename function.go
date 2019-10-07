package main

import "fmt"

// 関数を返す関数
func funcResultFunc() func() {
	n := 10
	return func() {
		fmt.Println(n)
	}
}

func main() {
	f := funcResultFunc()
	f()
}
