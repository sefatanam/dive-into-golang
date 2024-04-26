package strings

import "fmt"

/* Strings
Strings are immutable so you cannot change its content once created.
You need to first convert into a slice of rune
then do the change and in the end convert it back to string. */

/* Rune literals are just 32-bit integer values
(however they're untyped constants, so their type can change).
They represent unicode codepoints.
For example, the rune literal 'a' is actually the number 97. */

func UpdateString() {

	sampleStr := "Hello World"
	fmt.Println(sampleStr)
	fmt.Println(sampleStr[0])

	// update

	r := []rune(sampleStr)

	r[0] = 'W'

	sampleStr = string(r)

	fmt.Println(sampleStr)
}

func SwapRune(r rune) rune {
	switch {
	case 97 <= r && r <= 122:
		return r - 32
	case 65 <= r && r <= 90:
		return r + 32
	default:
		return r
	}
}
