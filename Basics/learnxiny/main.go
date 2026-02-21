package main

import (
	"fmt"
)

var limit = 10
var currentPosition = 3

func main() {
	fmt.Println(0)
	fmt.Println(1)
	fibonacci(0, 1)
}
func f(n int) {
  if n <= 1 {
		return n
	} else {
		return f(n-1) + f(n-2)
	}
}
func fibonacci(previous, next int) {
	if currentPosition > limit {
		return
	}

	currentNumber := previous + next
	currentPosition++
	previous = next
	next = currentNumber

	fmt.Println(currentNumber)
	fibonacci(previous, next)

}
