package main

import "fmt"

// typeキーワードによるインタフェース型の宣言
type myIF interface {
	typeName() string
	add(n int) int
}

// myIFインタフェースを実装する型の宣言(数値型)
type myInt int

// myInt型のtypeNameメソッドの宣言
func (m myInt) typeName() string {
	return "myInt"
}

// myInt型のaddメソッドの宣言
func (m myInt) add(n int) int {
	return int(m) + n
}

func main() {
	// myInt型の値をmyIF型の変数に設定
	var n myInt = 1
	var i myIF = n
	fmt.Println(i.typeName()) // myInt型のtypeNameメソッドを呼び出す
	fmt.Println(i.add(2))     // myInt型のaddメソッドを呼び出す
}
