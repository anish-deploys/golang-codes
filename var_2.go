package main

import (
	//if u do not want to use package but want to import it use _
	_ "bufio"
	"fmt"
)

func main2() {
	var a = "anish"
	fmt.Printf(a)

	fmt.Println()

	var b, c int = 1, 2
	fmt.Println(b, c)

	var e int
	fmt.Println(e)

	name := "anish"
	age := 20
	fmt.Printf("name : %s, age : %d\n", name, age)

	//shorthand dec can only be used once in a line
	g := 34
	h := 45

	fmt.Printf("g is %d, h is %d\n", g, h)

	//multiple var dec + ini in one line using shorthand dec
	as, ad := "anish", 58
	fmt.Printf("as is %s and ad is %d\n", as, ad)

	//if u do not want to use a var and want to dec then use _
	_ = 98

}
