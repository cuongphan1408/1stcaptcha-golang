package onestcaptcha

import "fmt"

type GoLike struct {
	value int
}

func NewGoLike(value int) *GoLike {
	return &GoLike{value}
}

func (g *GoLike) PrintValue() {
	fmt.Printf("Value: %d\n", g.value)
}

func (g *GoLike) Increment() {
	g.value++
}

func (g *GoLike) Decrement() {
	g.value--
}
