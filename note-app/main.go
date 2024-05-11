package main

import "fmt"

func main() {

	/* var array [5]int

	array[0] = 1
	array[2] = 2
	array[4] = 3
	array[3] = 45
	array[1] = 455

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	} */

	var twoDarray [2][3]int

	/* First row */
	twoDarray[0][0] = 1
	twoDarray[0][1] = 2
	twoDarray[0][2] = 3

	/* Second row */
	twoDarray[1][0] = 3
	twoDarray[1][1] = 4
	twoDarray[1][2] = 5

	rowLen := len(twoDarray)
	colLen := len(twoDarray[0])

	/* fmt.Println("ROW:", rowLen)
	fmt.Println("COL:", colLen) */

	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			fmt.Print(twoDarray[i][j], " ")
		}

		fmt.Print("\n")
	}

}
