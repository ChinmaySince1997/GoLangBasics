package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var output map[string]int
	output =  make(map[string]int)
	strArr := strings.Fields(s)
	for _, v := range strArr {
		_, ok := output[v]
		if (ok == true) {
			output[v] += 1
		} else {
			output[v] = 1
		}
	}
	return output
}

func main() {
	wc.Test(WordCount)
}
