package main

import "fmt"

// @solution-sync:begin
// 解题思路：
// 1. 以第一个字符串为基准，遍历字符串中的字符进行纵向对比
// 2. 如果出现字符串长度相等或者字符串中字符不等时返回结果
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) == i || strs[0][i] != strs[j][i] {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}

// @solution-sync:end

func main() {
	var strs = []string{"flower", "flow", "flight"}

	var result = longestCommonPrefix(strs)
	fmt.Printf("%v\n", result)
}
