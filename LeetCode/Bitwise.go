package main

import "fmt"

func main() {
	fmt.Println(isOdd(4))
}

func isOdd(number int) bool {
	return (number & 1) == 1
}
