package interfaces

import "fmt"

/*
	两个结构体可以比较的
 */

type Sayer interface {
	Say()
}

type Dog struct {
	Name string
}

func (d Dog) Say() {
	fmt.Println("Dog")
}
