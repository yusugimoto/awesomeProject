package pkg1

import (
	"awesomeProject/pkg2"
	"fmt"
)

type Foo struct {
	Name string
	name string
}

func Hello() {
	fmt.Println("Hello pkg1")
	pkg2.Hello()
}

func hello() {
	fmt.Println("can not call main")
	fmt.Println("call from pkg1 caller")
}

func PrintFoo(f Foo) {
	fmt.Println(f.Name)
	fmt.Println(f.name)
}