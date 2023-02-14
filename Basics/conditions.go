package main

import "fmt"

func add(a, b int) int  {
	
	defer fmt.Println("World")
	fmt.Println("Hello")
	return a + b 
}

// func main()  {
// 	// var test = true
// 	// if sum := add(10,5); sum < 10 {
// 	// 	fmt.Println("test is true", sum)
// 	// } else {
// 	// 	fmt.Println("test is false", sum)
// 	// }
// 	test := 
// 	switch {
// 	case test >= 1:
// 		fmt.Println("case 1")
// 	case test >= 2:
// 		fmt.Println("case 2")
// 	default:
// 		fmt.Println("default case")
// 	}
// }

func main()  {
	defer fmt.Println("World 1")
	defer fmt.Println("World 2")
	fmt.Println("Hello 1")
	add(1,2)
}