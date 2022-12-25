package main

import (
	"fmt"
	"math"
)

func main() {

	numbers := []int{2, 22, 33, 34, 35, 45, 56, 78, 89}
	fmt.Println(binarySearch(numbers, 35))
	/*numberForPivot := []int{1, 2, 3, 4, 5, 6, 7, 1, 2, 4}
	duplicateNumberForPivot := []int{1, 2, 3, 4, 5, 6, 6, 7, 7, 7, 8, 9, 9, 1, 2, 4}
	fmt.Println("Search index is => ", binarySearch(numbers, 78))
	fmt.Println("Pivot point is => ", findPivot(numberForPivot))
	fmt.Println("Pivot point is => ", findPivotInDuplicateSortedArray(duplicateNumberForPivot))*/

	//nestedNumber := [][]int{
	//	{1, 2, 3, 4, 5},
	//	{7, 8, 9, 10, 12},
	//	{13, 14, 15, 16, 17},
	//	{18, 19, 20, 21, 22},
	//}
	//
	//result := search2dArray(nestedNumber, 4)
	//fmt.Println("Search Index => ", result)
}

func binarySearch(nums []int, target int) int {
	startIndex := 0
	endIndex := len(nums) - 1
	for startIndex < endIndex {
		midIndex := startIndex + (endIndex-startIndex)/2

		if nums[midIndex] < target {
			startIndex = midIndex + 1
		} else if nums[midIndex] > target {
			endIndex = midIndex - 1
		} else {
			return midIndex
		}
	}
	return -1
}

func findPivot(nums []int) int {
	startIndex := 0
	endIndex := len(nums) - 1

	for startIndex < endIndex {
		midIndex := startIndex + (endIndex-startIndex)/2
		if nums[midIndex] > nums[midIndex+1] {
			return midIndex
		} else if nums[midIndex-1] > nums[midIndex] {
			return midIndex - 1
		} else if nums[midIndex] >= nums[startIndex] {
			startIndex = midIndex
		} else {
			endIndex = midIndex - 1
		}
	}
	return -1
}

func findPivotInDuplicateSortedArray(nums []int) int {
	startIndex := 0
	endIndex := len(nums) - 1
	for startIndex < endIndex {
		midIndex := startIndex + (endIndex-startIndex)/2
		if nums[startIndex] == nums[midIndex] && nums[midIndex] == nums[endIndex] {
			if nums[startIndex] > nums[startIndex+1] {
				return startIndex
			}
			startIndex++
			if nums[endIndex] > nums[endIndex] {
				return endIndex
			}
			endIndex++
		} else if nums[midIndex] > nums[midIndex+1] {
			return midIndex
		} else if nums[midIndex-1] > nums[midIndex] {
			return midIndex - 1
		} else if nums[midIndex] >= nums[startIndex] {
			startIndex = midIndex
		} else {
			endIndex = midIndex - 1
		}
	}
	return -1
}

func rotatedCount(nums []int) int {
	return findPivot(nums) + 1
}

func checkArrayIsSortedOrRotated(nums []int) bool {
	rotatedCount := 0
	length := len(nums)
	for i := 0; i < length; i++ {

		if nums[i] > nums[(i+1)%length] {
			rotatedCount++
		}
	}

	if rotatedCount > 1 {
		return false
	} else {
		return true
	}

}

func splitArray(nums []int, m int) int {
	startNum := 0
	endNum := 0

	for _, num := range nums {
		startNum = int(math.Max(float64(num), float64(startNum)))
		endNum += num
	}

	for startNum < endNum {
		midNum := startNum + (endNum-startNum)/2
		subArraySum := 0
		pieces := 0

		for _, num := range nums {
			if subArraySum+num > midNum {
				pieces++
				subArraySum = num
			} else {
				subArraySum += num
			}
		}

		if pieces > m {
			startNum = midNum + 1
		} else {
			endNum = midNum
		}

	}

	return endNum

}

func search2dArray(nums [][]int, target int) bool {

	row := 0
	col := len(nums[0]) - 1

	for row < len(nums) && col >= 0 {
		if col == 0 {
			if nums[row][col] == target {
				return true
			}
		} else if nums[row][col] == target {
			return true
		} else if nums[row][col] < target {
			row++
		} else {
			col--
		}

	}
	return false

}
