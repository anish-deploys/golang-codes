package main

import (
	"fmt"
)

func main5() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("range", i)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	//There is no explicit while loop here, for can be used as traditional, while and infinite loop

	//looping over collections
	//collections - array, slice, map, struct, channel
	numbers := []int{1, 2, 3}
	for index, value := range numbers {
		//range is used loop over slices
		fmt.Println(index, value)
	}

	x, y := 1, 2 //this is ok if y is new
	fmt.Println(x + y)
}
