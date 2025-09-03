package main

import (
	"fmt"
	"slices"
)

func main9() {

	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)
	//even a non-nil slice can have a length of 0.

	s = make([]string, 3)
	//Creates a slice of 3 strings, all set to the zero value ("")
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	//created an empty slice c of the same length as s and copy into c from s
	copy(c, s)
	fmt.Println("cpy:", c)

	//slice slicing
	//Slices support a “slice” operator with the syntax slice[low:high]
	l := s[2:5] //index 2 is inclusive and 5 is exclusive
	fmt.Println("sl1:", l)

	l = s[:5] //from beginning to index 5
	fmt.Println("sl2:", l)

	l = s[2:] // From index 2 to end
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	//The length of the inner slices can vary, unlike with multi-dimensional arrays
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1 //For each row i, we create a slice of integers of length i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
