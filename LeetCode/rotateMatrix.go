package main

import "fmt"

func printArray(arr [][]int) {
	for _, row := range arr {
		fmt.Println(row)
	}
}
func main() {
	twoDArray := [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}}

	// Print the initial array
	fmt.Println("Initial array:")
	printArray(twoDArray)

	// Rotate the array
	rotateArray(&twoDArray)

	// Print the rotated array
	fmt.Println("Rotated array:")
	printArray(twoDArray)
}

func rotateArray(arr *[][]int) {

	// Get the number of rows and columns in the array
	rows := len(*arr) - 1
	cols := len((*arr)[0]) - 1

	// Swap each element in the array individually
	for i := 0; i < rows/2; i++ {
		for j := i; j < cols-i; j++ {
			temp := (*arr)[i][j]
			/*Top by left*/
			(*arr)[i][j] = (*arr)[rows-j][i]
			/*left by bottom*/
			(*arr)[rows-j][i] = (*arr)[rows-i][cols-j]
			/*bottom by right*/
			(*arr)[rows-i][cols-j] = (*arr)[j][cols-i]
			/*right by top */
			(*arr)[j][cols-i] = temp

			fmt.Println("The J =", j)
			printArray((*arr))
		}
		fmt.Println("The I completed:", i)
		printArray((*arr))
	}
}
