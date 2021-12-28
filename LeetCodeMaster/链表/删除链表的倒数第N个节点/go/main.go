package main


import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func Remove(head *ListNode, num int) *ListNode {
	if head == nil {
		return nil
	}

	dumyHead := &ListNode{Val:0, Next:head}
	cur := dumyHead
	fast := dumyHead
	for i:= 0; i < num;i++ {
		fast = fast.Next
		fmt.Printf("i = %d\n", i)
	}

	for fast.Next != nil {
		cur = cur.Next
		fast = fast.Next
	}

	cur.Next = cur.Next.Next
	return dumyHead.Next
}

func main() {
	node5 := &ListNode{5, nil}
	node4 := &ListNode{4, node5}
	node3 := &ListNode{3, node4}
	node2 := &ListNode{2, node3}
	node1 := &ListNode{1, node2}
	head := Remove(node1, 1)

	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Printf("\n")
}