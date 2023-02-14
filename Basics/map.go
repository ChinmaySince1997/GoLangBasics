package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main()  {
	var m map[string]Vertex
	fmt.Println(m, len(m));
	m = make(map[string]Vertex)
	fmt.Println(m, len(m));
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"], m, len(m));	
	var n = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	
	fmt.Println(n, len(n), n["Google"], n["google"])

	k := map[string]Vertex{
		"Chinmay": {
			10.932,32.222,
		},
	}
	fmt.Println(k, len(k), k["Chinmay"], k["google"])
}