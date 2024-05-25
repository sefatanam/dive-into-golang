package articles

import "fmt"

func RunArticles() {

	i, j := 42, 64
	fmt.Println("Initial Value (i,j) :=", i, j)

	p := &i // point to i
	fmt.Println("P point to :=", p)
	fmt.Println("P value is :=", *p)

	*p = 32
	fmt.Println("P value is :=", *p)
	fmt.Println("Initial Value (i,j) :=", i, j)

	// Why

	fmt.Println("P point to :=", p)
	fmt.Println("I point to :=", &i)

}
