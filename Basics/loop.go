package main

import (
	"fmt"
	"time"
)

func main()  {
	// sum := 0
	// for i := 0; i < 10; i++ {
	// 	sum += i	
	// }
	sum := 0
	for sum < 10 {
		sum += 1
	}
	fmt.Printf("Sum: %v\n", sum)
}