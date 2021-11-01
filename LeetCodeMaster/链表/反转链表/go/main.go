package main

import (
	"fmt"
)
type Node struct {
	Val int
	Next *Node
}
func Reverse(head *Node) *Node {
	if head == nil {
		return nil
	}

	var cur,pre,tmp *Node
	cur = head
	pre = nil
	tmp = nil
	for cur != nil {
		tmp = cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	return pre
}

func main() {
	node5 := &Node{5, nil}
	node4 := &Node{4, node5}
	node3 := &Node{3, node4}
	node2 := &Node{2, node3}
	node1 := &Node{1, node2}
	head := Reverse(node1)

	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Printf("\n")
}