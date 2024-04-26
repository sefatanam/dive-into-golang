package main

import "fmt"

func main() {
	result := add(12, 12.12)
	fmt.Println(result)
}

type MinType interface {
	int | float32 | float64
}

func add[T MinType](a T, b T) T {
	return a + b
}
