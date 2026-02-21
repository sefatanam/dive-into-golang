package main

import "fmt"

func main() {
	str := [6]string{"this", "is", "a", "go", "interview", "question"}

	sl1 := str[1:4] // Len = 3, Capacity = 5

	fmt.Println("Slice 1 :", "Len ", len(sl1), "Capacity ", cap(sl1))

	sl2 := sl1[2:3]  // Len = 1, Capacity = 3 
	fmt.Println("Slice 2 :", "Len ", len(sl2), "Capacity ", cap(sl2))
}
