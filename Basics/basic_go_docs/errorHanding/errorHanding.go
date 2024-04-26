package errorhanding

import "fmt"

func ErrorMain() {
	defer safeExist()
	const one = 2
	if one != 1 {
		panic("One is not equal to 1.")
	}
}

func safeExist() {
	if r := recover(); r != nil {
		fmt.Println("Panic recovered!!")
	}
}
