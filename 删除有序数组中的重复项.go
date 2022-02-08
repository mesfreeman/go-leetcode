package main

import "fmt"

// @solution-sync:begin
func removeDuplicates(nums []int) int {
	// 如果没有元素，直接返回
	n := len(nums)
	if n == 0 {
		return 0
	}

	s := 1                   // 定义一个慢指针，用于记录不同元素的填入的下标位置
	for f := 1; f < n; f++ { // 遍历切片，如果元素不同，则填入相应位置中
		if nums[f-1] != nums[f] {
			nums[s] = nums[f]
			s++
		}
	}
	return s
}

// @solution-sync:end

/*
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

输入：nums = [0,0,1,1,1,2,2,3,3,4]
输出：5, nums = [0,1,2,3,4]
解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。
*/
func main() {
	var nums = []int{1, 1, 2}

	var result = removeDuplicates(nums)
	fmt.Printf("%v\n", result)
}
