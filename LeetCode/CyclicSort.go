package main

import "fmt"

func main() {
	//	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	//nums2 := []int{4, 3, 2, 7, 8, 2, 3, 1}

	// fmt.Println(missingNumber(nums))
	// fmt.Println(nums2)

	//fmt.Println(findAllMissingNumber(nums2))
	fmt.Println(findAllDuplicate([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

func missingNumber(nums []int) int {

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
			return i
		}
	}
	return len(nums)
}

func findAllMissingNumber(nums []int) []int {
	var result []int
	index := 0
	for index < len(nums) {
		correctIndex := nums[index] - 1
		if nums[index] != nums[correctIndex] {
			swap(nums, index, correctIndex)
		} else {
			index++
		}
	}

	for i, el := range nums {
		if el != i+1 {
			result = append(result, i+1)
		}
	}

	return result
}

// 1,3,4,2,2 => 2
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

func swap(nums []int, fromIndex int, toIndex int) {
	temp := nums[fromIndex]
	nums[fromIndex] = nums[toIndex]
	nums[toIndex] = temp
}
