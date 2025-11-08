package main

import "fmt"

func sum(nums ...int) {
	//function that will take an arbitrary number of ints as argument including none
	//... = ellipsis
	//useful when no of inputs may vary
	fmt.Print(nums, " ")
	//nums is treated as a slice of int inside the function
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main13() {

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
