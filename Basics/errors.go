package main

import (
	"fmt"
	"time"
)

type myError struct {
	occurringTime time.Time
	msg           string
}

func (e *myError) Error() string {
	return fmt.Sprintf("error Message: %v occurred at: %v", e.msg, e.occurringTime)
}

func run() error {
	return &myError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
	// i, err := strconv.Atoi("")
	// if err != nil {
	// 	fmt.Printf("couldn't convert number: %v\n", err)
	// 	return
	// }
	// fmt.Println("Converted integer:", i)
}
