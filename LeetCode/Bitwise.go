package main

import "fmt"

func main() {
	fmt.Println(isOdd(4))
	fmt.Println(isOdd(5))
	fmt.Println(isOdd(6))
}

/**
0 & 0 => 0 ( Even )
1 & 0 => 0 ( Even )
0 & 1 => 0 ( Even )
1 & 1 => 1 ( Odd )

EX - number 4
Scale : 8421
4 in binary => 100 & 1 => 100 => 4 ( Even )

EX - number 5
Scale : 8421
5 in binary => 101 & 1 => 111 => 7 ( Odd )
*/

func isOdd(number int) bool {
	res := number & 1
	fmt.Println(res)
	return (number & 1) == 1
}
