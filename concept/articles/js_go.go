package articles

import "fmt"

func RunArticles() {

	i, j := 42, 64
	fmt.Println("Initial Value (i,j) :=", i, j)

	p := &i // point to i
	fmt.Println("P point to :=", p)
	fmt.Println("P value is :=", *p)

	// *p = 32
	*p = PassByValue(p)
	fmt.Println("P value is :=", *p)
	fmt.Println("Initial Value (i,j) :=", i, j)

	// Why

	fmt.Println("P point to :=", p)
	fmt.Println("P point to with value:=", *p)
	fmt.Println("I point to :=", &i)
	fmt.Println("I point to with value :=", i)

}

func PassByRef(num *int) {
	*num = 54
}

func PassByValue(num *int) int {
	res := *num
	res += 10
	return res
}
