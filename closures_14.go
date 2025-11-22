package main

import "fmt"

func makeCounter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

func main14() {
	counter1 := makeCounter()
	fmt.Println(counter1()) // 1
	fmt.Println(counter1()) // 2
	fmt.Println(counter1()) // 3

	counter2 := makeCounter()
	fmt.Println(counter2()) // 1
	//counter1() increases its own count: 1 → 2 → 3
	//counter2() is a new counter, so it starts from 1
	//It remembers and modifies count every time it's called
	//Even though makeCounter() has finished, its local variable count still lives on inside the closure
}
