package main

import (
	"fmt"
	"math"
)

// @solution-sync:begin
func reverse(x int) int {
	res := 0
	for x != 0 {
		if res < math.MinInt32/10 || res > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10      // 先求余拿到末尾数
		x /= 10              // 然后将整数缩小10倍
		res = res*10 + digit // 将结果值放大10倍后加上末尾数
	}
	return res
}

// @solution-sync:end

/*
给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。

假设环境不允许存储 64 位整数（有符号或无符号）。

示例 1：
输入：x = 123
输出：321
*/
func main() {
	var x = 123

	var result = reverse(x)
	fmt.Printf("%v\n", result)
}
