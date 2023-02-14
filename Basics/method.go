package main

import "fmt"

type Vertex struct {
	x, y int
}

type VertexInterface interface {
	Add() int
	ChangeValue(int, int)
}

func (v *Vertex) Add() int {
	if v == nil {
		return 0
	}
	return v.x + v.y
}

func (v *Vertex) ChangeValue(x, y int) {
	if v == nil {
		return
	}
	v.x = x
	v.y = y
}

// func ChangeValue(v Vertex, x, y int) {
// 	v.x = x
// 	v.y = y
// }

func main() {
	var v VertexInterface = &Vertex{1, 2}
	fmt.Println(v.Add(), v)
	v.ChangeValue(10, 12)
	fmt.Println(v.Add(), v)
}
