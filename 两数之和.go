package main

import "fmt"

// @solution-sync:begin
func twoSum(nums []int, target int) []int {
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// @solution-sync:end

func main() {
	var nums = []int{2, 7, 11, 15}
	var target = 9

	var result = twoSum(nums, target)
	fmt.Printf("%v\n", result)
}
