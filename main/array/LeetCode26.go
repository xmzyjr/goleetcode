package main

import "fmt"

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	index1 := 0
	for i := 1; i < length; i++ {
		if nums[i] == nums[index1] {
			continue
		}
		index1++
		nums[index1] = nums[i]
	}
	return index1 + 1
}

func main() {
	nums1 := []int{1, 1, 2}
	duplicates := removeDuplicates(nums1)
	fmt.Println(nums1, duplicates)
}
