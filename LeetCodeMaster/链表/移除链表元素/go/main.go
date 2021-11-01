package main

import (
	"fmt"
)

type Node struct {
	Val int
	Next *Node
}

func RemoveVal(head *Node, target int) *Node {
	tmp := &Node {
		Next:head,
	}

	first := tmp
	second := tmp.Next

	for second != nil {
		if second.Val == target {
			second = second.Next
			first.Next = second
		} else {
			first = first.Next
			second = second.Next
		}
	}

	return tmp.Next
}

func RemoveVal2(head *Node, target int) *Node {
	if nil == head {
		return nil
	}

	for head != nil && head.Val == target {
		head = head.Next
	}

	if head == nil {
		return nil
	}

	first := head
	second := head.Next

	for second != nil {
		if second.Val == target {
			second = second.Next
			first.Next = second
		} else {
			first = first.Next
			second = second.Next
		}
	}

	return head
}

//1->2->3->4->5->6->nil

func main() {
	node1 := &Node {
		Val:5,
	}

	node2 := &Node {
		Val:5,
	}

	node3 := &Node {
		Val:5,
	}

	node4 := &Node {
		Val:5,
	}

	node5 := &Node {
		Val:5,
	}

	node6 := &Node {
		Val:5,
		Next:nil,
	}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = nil

	head := RemoveVal2(node1, 5)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}