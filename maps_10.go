package main

import (
	"fmt"
	"maps"
)

func main10() {
	//maps - similar to dictionaries in Python or objects in JS
	m := make(map[string]int)
	//make(map[key-type]val-type)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	//If the key doesn’t exist, the zero value(nil) of the value type is returned
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))

	//removes key/value pairs from a map
	delete(m, "k2")
	fmt.Println("map:", m)

	//To remove all key/value pairs from a map
	clear(m)
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// can also declare and initialize a new map in the same line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
