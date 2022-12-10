package main

import (
	"fmt"
)

func main() {
	//nums := []int{9, 0, 4, 2, 3, 5, 7, 0, 1}
	//fmt.Println(findAllMissingNumber(nums))

	nums := []int{3, -1, 4, 1, -1}
	res := firstMissingPositive(nums)
	fmt.Println(res)
}

/*
Main base of Cyclic sort [ T(O(n), M(O(1) ]
*/
func cyclicSort(nums []int) {
	index := 0
	for index < len(nums) {
		correctIndex := nums[index] - 1
		if nums[correctIndex] != nums[index] {
			swap(nums, index, correctIndex)
		} else {
			index++
		}
	}
}

/*
*
Find missing number using cyclic sort
*/
func missingNumber(nums []int) int {
	index := 0
	for index < len(nums) {
		correctIndex := nums[index] // As if element can be 0 or missing so current element is correct index
		if nums[index] < len(nums) && nums[correctIndex] != nums[index] {
			swap(nums, index, correctIndex)
		} else {
			index++
		}
	}
	for i, el := range nums {
		if i != el {
			return i
		}
	}
	return len(nums)
}

/*
*
Find all missing number using cycle sort
*/
func findAllMissingNumber(nums []int) []int {
	var result []int
	index := 0
	for index < len(nums) {
		correctIndex := nums[index]
		if nums[index] < len(nums) && nums[index] != nums[correctIndex] {
			swap(nums, index, correctIndex)
		} else {
			index++
		}
	}

	for i, el := range nums {
		if el != i {
			result = append(result, i)
		}
	}

	return result
}

// 4,3,1,2,2 => 2
func findDuplicate(nums []int) int {
	index := 0
	for index < len(nums) {
		if nums[index] != index+1 {
			correctIndex := nums[index] - 1
			if nums[index] != nums[correctIndex] {
				swap(nums, index, correctIndex)
			} else {
				return nums[index]
			}
		} else {
			index++
		}
	}
	return -1

}

// 4,3,2,7,8,2,3,1 => 2,3
func findAllDuplicate(nums []int) []int {
	var result []int
	index := 0
	for index < len(nums) {
		correctIndex := nums[index] - 1
		if nums[index] < len(nums) && nums[index] != nums[correctIndex] {
			swap(nums, index, correctIndex)
		} else {
			index++
		}
	}

	for i, el := range nums {
		if el != i+1 {
			result = append(result, el)
		}
	}

	return result

}

/*
*
Find errors of nums [repeated number , missing number]
1,2,2,4,5 => [2,3]
*/
func findErrorNums(nums []int) []int {
	var result []int
	index := 0
	for index < len(nums) {
		correctIndex := nums[index] - 1
		if nums[correctIndex] != nums[index] {
			swap(nums, index, correctIndex)
		} else {
			index++
		}
	}

	for i, num := range nums {
		if num != i+1 {
			result = append(result, nums[i], i+1)
			return result
		}
	}

	return result
}

/*
*
First missing positive
*/
func firstMissingPositive(nums []int) int {
	index := 0
	for index < len(nums) {
		correctIndex := nums[index] - 1 // As if element can be 0 or missing so current element is correct index
		if nums[index] > 0 && nums[index] <= len(nums) && nums[correctIndex] != nums[index] {
			swap(nums, index, correctIndex)
		} else {
			index++
		}
	}
	for i, el := range nums {
		if 1+i != el {
			return 1 + i
		}
	}
	return 1 + len(nums)
}

/*
*
Swap two number in array
*/
func swap(nums []int, fromIndex int, toIndex int) {
	temp := nums[fromIndex]
	nums[fromIndex] = nums[toIndex]
	nums[toIndex] = temp
}
