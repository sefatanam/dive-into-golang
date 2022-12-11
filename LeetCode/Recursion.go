package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := binarySearchr(nums, 7, 0, len(nums))

	fmt.Println(result)
}

func binarySearchr(nums []int, target int, startIndex int, endIndex int) int {

	if startIndex > endIndex {
		return -1
	}

	midIndex := startIndex + (endIndex-startIndex)/2

	if target == nums[midIndex] {
		return midIndex
	}

	if target < nums[midIndex] {
		return binarySearchr(nums, target, startIndex, midIndex-1)
	}
	return binarySearchr(nums, target, midIndex+1, endIndex)
}
