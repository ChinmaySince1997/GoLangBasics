package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main()  {
	v := []Vertex{
		{1,1},
		{2,2},
		{3,3},
	}
	fmt.Println(len(v), cap(v))
	// p := &v
	// p.X = 20
	// fmt.Println(p, v, &p, &v, *p, p.X, v.X)
}