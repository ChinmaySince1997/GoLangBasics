package greetings

import "fmt"

func Hello(name string) (message string) {
	message = fmt.Sprintf("Hi, %v. Welcome!", name)
	return
}
