package main

import "fmt"

func main() {
	twoDArray := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}

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
	/*// Get the number of rows and columns in the array
	rows := len(arr)
	cols := len(arr[0])

	// Create a new array with the rows and columns swapped
	rotated := make([][]int, cols)
	for i := range rotated {
		rotated[i] = make([]int, rows)
		for j := range rotated[i] {
			rotated[i][j] = arr[rows-j-1][i]
		}
	}
	return rotated*/

	/*	rows := len(*arr)
		cols := len((*arr)[0])

		for i := 0; i < rows/2; i++ {

			for j := 0; j < cols-i-1; j++ {
				temp := (*arr)[i][j]
				(*arr)[i][j] = (*arr)[rows-j-1][i]
				(*arr)[rows-j-1][i] = (*arr)[rows-i-1][cols-j-1]
				(*arr)[rows-i-1][cols-j-1] = (*arr)[j][cols-i-1]
				(*arr)[j][cols-i-1] = temp

				fmt.Println("The J =", j)
				printArray((*arr))

			}
			fmt.Println("The I completed:", i)
			printArray((*arr))
		}*/

	// Get the number of rows and columns in the array
	rows := len(*arr)
	cols := len((*arr)[0])

	// Swap each element in the array individually
	for i := 0; i < rows/2; i++ {
		for j := i; j < cols-i-1; j++ {
			temp := (*arr)[i][j]
			/*Top by left*/
			(*arr)[i][j] = (*arr)[rows-j-1][i]
			/*left by bottom*/
			(*arr)[rows-j-1][i] = (*arr)[rows-i-1][cols-j-1]
			/*bottom by right*/
			(*arr)[rows-i-1][cols-j-1] = (*arr)[j][cols-i-1]
			/*right by top */
			(*arr)[j][cols-i-1] = temp

			fmt.Println("The J =", j)
			printArray((*arr))
		}
		fmt.Println("The I completed:", i)
		printArray((*arr))
	}
}

func printArray(arr [][]int) {
	for _, row := range arr {
		fmt.Println(row)
	}
}
