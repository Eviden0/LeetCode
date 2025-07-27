package code

import "sort"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// 1. 遍历链表，将值存入切片
	tnums := make([]int, 0)
	for node := head; node != nil; node = node.Next {
		tnums = append(tnums, node.Val)
	}

	// 2. 对切片排序
	sort.Ints(tnums)

	// 3. 将排序后的值重新赋值到链表
	node := head
	for _, val := range tnums {
		node.Val = val
		node = node.Next
	}

	// 4. 返回链表头
	return head
}
