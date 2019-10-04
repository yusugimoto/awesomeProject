package pkg1

import "fmt"

type Foo struct {
	Name string
	name string
}

func Hello() {
	fmt.Println("Hello pkg1")
}

func hello() {
	fmt.Println("can not call main")
	fmt.Println("call from pkg1 caller")
}

func PrintFoo(f Foo) {
	fmt.Println(f.Name)
	fmt.Println(f.name)
}