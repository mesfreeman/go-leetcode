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
// 解题思路：
// 1. 根据回文数特点，负数不可能，判断后提前返回
// 2. 将数字转为字符串，然后对字符串进行遍历对比
func isPalindrome02(x int) bool {
	if x < 0 {
		return false
	}

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

func main() {
	var x = 121

	var result = isPalindrome02(x)
	fmt.Printf("%v\n", result)
}
