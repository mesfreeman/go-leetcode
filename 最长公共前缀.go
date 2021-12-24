package main

import "fmt"

// @solution-sync:begin
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 以第一个字符串为基准，遍历字符串中的字符进行纵向对比
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			// 如果出现字符串长度相等或者字符串中字符不等时返回结果
			if len(strs[j]) == i || strs[0][i] != strs[j][i] {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}

// @solution-sync:end

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1：
输入：strs = ["flower","flow","flight"]
输出："fl"
*/
func main() {
	var strs = []string{"flower", "flow", "flight"}

	var result = longestCommonPrefix(strs)
	fmt.Printf("%v\n", result)
}
