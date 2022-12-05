package main

import (
	"fmt"
)

func main() {

	numbers := []int{2, 22, 33, 34, 35, 45, 56, 78, 89}
	numberForPivot := []int{1, 2, 3, 4, 5, 6, 7, 1, 2, 4}
	duplicateNumberForPivot := []int{1, 2, 3, 4, 5, 6, 6, 7, 7, 7, 8, 9, 9, 1, 2, 4}
	fmt.Println("Search index is => ", binarySearch(numbers, 78))
	fmt.Println("Pivot point is => ", findPivot(numberForPivot))
	fmt.Println("Pivot point is => ", findPivotInDuplicateSortedArray(duplicateNumberForPivot))
}

func binarySearch(nums []int, target int) int {
	startIndex := 0
	endIndex := len(nums) - 1
	for startIndex < endIndex {
		midIndex := startIndex + (endIndex-startIndex)/2

		if nums[startIndex] < target {
			startIndex = midIndex + 1
		} else if nums[startIndex] > target {
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
