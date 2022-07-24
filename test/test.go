package main

import "fmt"

type P struct {
	name string
	age  int
}

type w struct {
	P
}
type r struct {
	P
}

func (w w) use() {
	fmt.Println("Дерево")
}
func (r r) use() {
	fmt.Println("Камень")
}

type I interface {
	use()
}
type BI struct {
	I
}

//
func main() {
	build := BI{I: w{P{name: "ivan", age: 22}}}
	build.use()
	builds := BI{I: r{P{name: "petro", age: 22}}}
	builds.use()
}
