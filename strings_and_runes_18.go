package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 1. Basic string and length
	s := "Go語" // "語" is a 3-byte Unicode character
	fmt.Println("Original string:", s)
	fmt.Println("Byte length:", len(s))                   // 6 bytes
	fmt.Println("Rune count:", utf8.RuneCountInString(s)) // 3 characters

	fmt.Println()

	// 2. Printing raw UTF-8 bytes
	fmt.Println("Bytes (hex):")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i]) // print each byte
	}
	fmt.Println()

	fmt.Println()

	// 3. Iterating with range (gives runes, not bytes)
	fmt.Println("Runes (characters):")
	for i, r := range s {
		fmt.Printf("Index %d: %c (Unicode: %U)\n", i, r, r)
	}

	fmt.Println()

	// 4. Convert string to []rune to modify
	runes := []rune(s)
	runes[1] = 'O' // replace '語' with 'O'
	modified := string(runes)
	fmt.Println("Modified string:", modified)

	fmt.Println()

	// 5. Concatenation
	part1 := "Hello"
	part2 := "世界"
	combined := part1 + ", " + part2 + "!"
	fmt.Println("Concatenated string:", combined)

	fmt.Println()

	// 6. Convert string to []byte and back
	b := []byte(s)
	fmt.Println("String to bytes:", b)
	sFromBytes := string(b)
	fmt.Println("Bytes to string:", sFromBytes)

	// 7. Accessing the N-th Character:
	str := "Hello"
	runes1 := []rune(str)
	fmt.Println(string(runes1[2])) //prints 3rd char

	//8. Changing a Character in a String
	str1 := "Hello"
	runes2 := []rune(str1)
	runes2[0] = 'J'
	fmt.Println(string(runes2))

}
