package main

import "fmt"

// @solution-sync:begin
// 解题思路：
// 1. 根据回文数，负数不可能，判断后提前返回
// 2. 对数字进行求余操作，将个个数字放入切片中，然后对切片遍历对比
func isPalindrome(x int) bool {
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

func main() {
	var x = 121

	var result = isPalindrome(x)
	fmt.Printf("%v\n", result)
}
