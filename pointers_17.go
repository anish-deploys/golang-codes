package main

import (
	"fmt"
)

func changeValue(num *int) {
	*num = 42
}
func main17() {
	x := 10
	fmt.Println(x)

	changeValue(&x)
	fmt.Println(x)
	fmt.Println(&x)
}
