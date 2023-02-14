package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)\n", p.Name, p.Age)
}

type Vertex struct {
	X, Y int
}

func (p Vertex) String() string {
	p.X = 30
	return fmt.Sprintf("%v (%v years)\n", p.X, p.Y)
}

func main() {
	x := Vertex{10, 10}
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z, x, x.String())
}
