package main

import "fmt"

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

func NewList(nums []int, pos int) *ListNode {
	head := &ListNode{}
	cur := head
	for _, num := range nums {
		cur.Next = &ListNode{Val: num}
		cur = cur.Next
	}
	temp := head
	for i := 0; i < pos+1; i++ {
		temp = temp.Next
	}
	cur.Next = temp
	return head.Next
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}

func main() {
	//"n2"
	//"n2logn"
	//"n3"
	nums := []int{3, 2, 0, -4}
	pos := 1
	//var head *ListNode
	head := NewList(nums, pos)
	p := detectCycle(head)
	fmt.Println(p)

}
