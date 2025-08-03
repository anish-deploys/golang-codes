package main

import (
	_ "bufio"
	"fmt"
)

// const must be defined at the package/fxn level to use, it cant be dynamically used
// cannot be assigned runtime values like fxn calls or vars
const e = 2.71

func main4() {
	const pi = 3.1415
	fmt.Println(pi)
	//now this const cannot be updated like var

	const p, yt = 98, "ytt"
	fmt.Print(p, yt, "\n")

	a := 90
	fmt.Println(a)

	//typed const
	const num int = 42

	//untyped const
	const num1 = 49
	var _ float64 = num1

	fmt.Println(num, num1)

	/* const must be computed at compile time
	const x = math.Sqrt(25) //not allowed
	*/

	//iota
	//used to create incrementing constants automatically

	const (
		a1 = iota //0
		a2        //1
		a3        //2
	)
	fmt.Println(a1, a2, a3)

	//bitwise op with iota
	//Used in permissions, flags, and masks.
	const (
		read    = 1 << iota // 1 (0001)
		write               // 2 (0010)
		execute             // 4 (0100)
	)

	fmt.Println(read, write, execute)

}
