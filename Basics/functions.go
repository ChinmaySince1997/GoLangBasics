package main

import (
	"fmt"
)

// func add(a int, b int) int  {
// 	return a + b 
// }

func add(a, b int) int  {
	return a + b 
}

// func swap(a, b string) (string, string)  {
// 	return b, a
// }

func swap(a, b string) (x, y string)  {
	x = b
	y = a
	return
}

func main()  {
	// fmt.Println(add(3, 5))
	fmt.Println(swap("Hello", "World"))
}