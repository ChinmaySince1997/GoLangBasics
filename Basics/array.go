package main

import "fmt"

func main()  {
	// var arr [2]string
	// arr[0] = "Chinmay"
	// arr[1] = "Kiran"
	// fmt.Println(arr, arr[0], arr[1])
	// primes := [6]int{2, 3, 5, 7, 11, 13}
	// var s []int = primes[:]
	// primes[0] =  1
	// fmt.Println(s)
	// fmt.Println(primes)

	// names := make([]string, 4, 4)
	// names[0] = "John"
	// names[1] = "Paul"
	// names[2] = "George"
	// names[3] = "Ringo"
	 names :=[]string {
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	cpNames := append(names, "test", "well")
	// names[0] = "new"
	fmt.Println(cpNames, cap(cpNames), len(cpNames), names)
	
	// numsCopy := names
	// a := names[0:2]
	// b := names[1:3]
	// fmt.Println(a, b)

	// b[0] = "XXX"
	// fmt.Println(a, b)
	// fmt.Println(names, numsCopy)

	for i, v := range cpNames {
		fmt.Println(v, "  ", i)
	}
}