package main

import "fmt"

type hogefuga struct {
	int
	hogehoge // 構造体の中に構造体を持つことができる。この場合名前をつけなくても良い
	fugafuga // 名前をつけない＝匿名フィールド
	pi piyopiyo // 匿名フィールドではない
}

type hogehoge struct {
	name string
	x int
}

type fugafuga struct {
	name string
	y int
}

type piyopiyo struct {
	name string
	z int
}

func main() {
	var hf hogefuga
	hf.int = 1000 // 「構造体名.型名」でアクセスできる
	hf.x = 1 // 構造体の匿名フィールドなのでhf.hogehoge.x をhf.xとしてアクセスできる
	hf.y = 2 // もちろんhf.fugafuga.y = 2でも代入することができる
	hf.pi.z = 20 //匿名フィールドではないのでhf.zとすることはできない
	hf.pi.name = "piyopiyo"
	hf.hogehoge.name = "hogehoge" // hogehoge, fugafuagaもnameをもっているのでhf.nameで代入することができない
	hf.fugafuga.name = "fugafuga"
	fmt.Print(hf)
}