package main

import "fmt"

func vals() (int, int) {
	//Go has built-in support for multiple return values, used often in idiomatic Go
	return 3, 7
}

func main12() {

	//If you only want a subset of the returned values, use the blank identifier _
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	//c gets the second return value (7)
	fmt.Println(c)
}
