package main

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// @solution-sync:begin
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	// 递归处理，如果表1小于表2，则以表1为起始
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}

	// 否则以表2为起始
	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2
}

// @solution-sync:end

/*
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
*/
func main() {
	var list1 = parseListNode([]int{1, 2, 4})
	var list2 = parseListNode([]int{1, 3, 4})

	var result = mergeTwoLists(list1, list2)
	fmt.Printf("%v\n", nodeToString(result))
}

func parseListNode(values []int) *ListNode {
	var head *ListNode
	var tail *ListNode
	for i := 0; i < len(values); i++ {
		var node = &ListNode{Val: values[i]}
		if head == nil {
			head = node
		} else {
			tail.Next = node
		}
		tail = node
	}
	return head
}

func nodeToString(h *ListNode) string {
	var buf strings.Builder
	buf.WriteString("[")
	var c = h
	for c != nil {
		buf.WriteString(fmt.Sprintf("%v", c.Val))
		if c.Next != nil {
			buf.WriteString(",")
		}
		c = c.Next
	}
	buf.WriteString("]")
	return buf.String()
}
