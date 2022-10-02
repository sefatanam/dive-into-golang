// Article link: https://golangdocs.com/pointers-in-golang
// SYNTAX OF POINTER => var ptr *type

package main

import "fmt"

// Pointer as a function argument
func f(a *int) {
	fmt.Println("print pointer name", a)
	fmt.Println("print pointer address", &a)
	fmt.Println("print pointer value", *a)
}

type Human struct {
	FirstName string
	LastName  string
	Age       int
}

func identity(h *Human) {
	fmt.Println("This human name is", *&h.FirstName, *&h.LastName, "and age is :", *&h.Age)
}

//  Returning a pointer from a function

func p() *int {
	v := 10
	return &v
}

func main() {
	fmt.Println("Pointer in Go")

	// var ptrint int =43
	// var qtrint *int
	// qtrint=&ptrint
	// fmt.Println("print pointer name",qtrint) // print pointer name 0xc0000a4018
	// fmt.Println("print pointer with [*]",*qtrint)// will print value of pointer
	// fmt.Println("print pointer with [&]",&qtrint) // will print address of pointer

	// human := Human{FirstName: "Rahim", LastName: "Uddin", Age: 12}

	// human2 := new(Human) //The new function in Go returns a pointer to a type.
	// human2.FirstName = "Karim"
	// human2.LastName = "Uddin"
	// human2.Age = 13

	// // Here is the differences between new and make
	// identity(&human)
	// identity(human2)

	// n := p()

	// fmt.Println(*n)

	//Pointers to a function

	f := func() {
		fmt.Println("Hello World")
	}
	pf := f
	pf()
}
