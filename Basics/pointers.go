package main

import "fmt"

func main()  {
	var i =  100
	p := &i
	// if (*p == i ) {
	// 	fmt.Println("is same:")
	// }
	fmt.Println("value:", i, p, &i, &p, *p)
}