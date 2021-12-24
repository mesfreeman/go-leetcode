package main

import (
	"fmt"
	"strconv"
)

// @solution-sync:begin
// 解题思路：
// 1. 根据回文数特点，负数不可能，判断后提前返回
// 2. 对数字进行求余操作，将个个数字放入切片中，然后对切片遍历对比
func isPalindrome01(x int) bool {
	if x < 0 {
		return false
	}

	var sl []int
	sl = make([]int, 0)
	for x != 0 {
		sl = append(sl, x%10)
		x /= 10
	}

	l := len(sl)
	for i := 0; i < l/2; i++ {
		if sl[i] != sl[l-i-1] {
			return false
		}
	}
	return true
}

// @solution-sync:end

// @solution-sync:begin
func isPalindrome02(x int) bool {
	// 根据回文数特点，负数不可能，判断后提前返回
	if x < 0 {
		return false
	}

	// 将数字转为字符串，然后对字符串进行遍历对比
	str := strconv.Itoa(x)
	l := len(str)
	for i := 0; i < l/2; i++ {
		if str[i] != str[l-i-1] {
			return false
		}
	}
	return true
}

// @solution-sync:end

/*
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。

示例 1：
输入：x = 121
输出：true
*/
func main() {
	var x = 121

	var result = isPalindrome02(x)
	fmt.Printf("%v\n", result)
}
