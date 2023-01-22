package main

import "fmt"

type person struct {
	name string
	age  int
}

var demoPersonList = []person{
	{age: 10, name: "Sefat Anam"},
	{age: 20, name: "Sefat Anam"},
	{age: 30, name: "Sefat Anam"},
}

func sum(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	/*for index, value := range demoPersonList {
		fmt.Printf("Index is: %v Value is %v\n", index, value)
	}*/

	/*Verdict Function */
	/*
		Variadic functions can be called with any number of trailing arguments.
		For example, fmt.Println is a common variadic function
	*/

	sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	sum(1, 2, 3, 4)
	sum(1, 2)

	/*Next Class https://gobyexample.com/closures*/

}
