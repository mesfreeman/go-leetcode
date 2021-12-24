package main

import "fmt"

// @solution-sync:begin
func twoSum(nums []int, target int) []int {
	l := len(nums)
	// 两层遍历，如果之和等于目标值，直接返回下标
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

/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。

示例 1：
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1]
*/
func main() {
	var nums = []int{2, 7, 11, 15}
	var target = 9

	var result = twoSum(nums, target)
	fmt.Printf("%v\n", result)
}
