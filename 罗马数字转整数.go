package main

import "fmt"

// @solution-sync:begin
// 解题思路：
// 罗马数字抽象为两种情况：
// 1）左边数字大于等于右边数字，此时值为数字相加之和
// 2）左边数字小于右边数字，此时值为右边数字减左边数字，也可以理解为左边的负值与右边数字之和
func romanToInt(s string) int {
	m := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	res := 0
	l := len(s)
	for i := 0; i < l; i++ {
		if i < l-1 && m[s[i]] < m[s[i+1]] {
			res -= m[s[i]]
		} else {
			res += m[s[i]]
		}
	}
	return res
}

// @solution-sync:end

func main() {
	var s = "III"

	var result = romanToInt(s)
	fmt.Printf("%v\n", result)
}
