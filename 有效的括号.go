package main

import "fmt"

// @solution-sync:begin
func isValid(s string) bool {
	n := len(s)

	// 如果字符串是奇数，肯定不是有效的括号，直接返回
	if n%2 == 1 {
		return false
	}

	// 定义有效括号映射，用于判断是左括号还是右括号
	pairs := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	var stacks []byte
	for _, v := range s {
		// 如果存在，则说明是右括号，否则是左括号
		if m, ok := pairs[byte(v)]; ok {
			// 判断切片里是否有数据，如果没有或者最后一个值不匹配则说明无效，直接返回
			if len(stacks) == 0 || stacks[len(stacks)-1] != m {
				return false
			}

			// 匹配到时，将该括号做出栈处理
			stacks = stacks[:len(stacks)-1]
		} else {
			// 左括号直接追加到切片中，做压栈处理
			stacks = append(stacks, byte(v))
		}
	}

	// 如果最后切片中的括号全部出栈，则表示是有效的括号
	return len(stacks) == 0
}

// @solution-sync:end

func main() {
	var s = "()"

	var result = isValid(s)
	fmt.Printf("%v\n", result)
}
