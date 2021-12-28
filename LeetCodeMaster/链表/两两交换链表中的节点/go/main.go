package main

import (
	"fmt"
)

type Node struct {
	Val int
	Next *Node
}

func twotwoSwap(head *Node) *Node {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	first := head
	second := head.Next
	var pre *Node

	tmpHead := second
	for first != nil && second != nil {
		tmp := second.Next

		first.Next = tmp
		second.Next = first
		if pre != nil {
			pre.Next = second
		}

		pre = first
		first = tmp
		if first != nil {
			second = first.Next
		} else {
			second = nil
		}
	}

	return tmpHead
}
func main() {
	node5 := &Node{5, nil}
	node4 := &Node{4, node5}
	node3 := &Node{3, node4}
	node2 := &Node{2, node3}
	node1 := &Node{1, node2}
	head := twotwoSwap(node1)

	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Printf("\n")
}