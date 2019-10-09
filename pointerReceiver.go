package main

import "fmt"

// 名前付きの型を作成
type myInt int

// myInt型にメソッドを追加
func (m myInt) addValue(n myInt) {
	m += n
}

// *myInt型にメソッドを追加
func (m *myInt) addPointer(n myInt) {
	fmt.Println("addPointer *m ",*m)
	*m += n
}

func main() {
	fmt.Println("---------value----------")
	// 値レシーバを持つメソッドの呼び出し
	var v myInt = 1
	// fmt.Println(*v) ERROR
	fmt.Println(&v)
	fmt.Println(v)
	v.addValue(1)
	v.addPointer(10)
	/*
		vは値なので ポインタ型intに対して追加されたmethodaddPointerを呼べないはずだが、
		実際には呼び出すことができる。これはコンパイラによるレシーバの暗黙的変換が行われ、(&v).addPointerと解釈されているためである
		個人的には型に厳格な言語?で型がなあなあになるような機能はいらない
	 */

	fmt.Println("---------pointer----------")
	var p = &v
	p.addValue(1) // コンパイラにより(*p).addValue(1)と暗黙的変換されている
	p.addPointer(10)
	fmt.Println(*p)
	fmt.Println(p)
	fmt.Println(&p)

}
