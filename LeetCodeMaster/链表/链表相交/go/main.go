package main

import (
	"fmt"
)

type ListNode  struct {
	Val int
	Next *ListNode
}

func FindCross2(headA * ListNode, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	
	curA := headA
	curB := headB

	stopA, stopB := false, false
	for {
		if curA == curB {
			return curA
		}

		if curA.Next == nil{
			if stopA == true && stopB == true {
				break
			}
			stopA = true
			curA = headB
			continue
		}

		if curB.Next == nil {
			if stopA == true && stopB == true {
				break
			}
			stopB = true
			curB = headA
			continue
		}

		curA = curA.Next
		curB = curB.Next
	}

	return nil
}

// 暴力方法，将A的地址都保存下来，然后遍历另一个链表，与保存的上个链表节点的地址逐一进行比对
func FindCross1(headA *ListNode, headB *ListNode) *ListNode{
	var nodeAddrA []*ListNode

	cur := headA
	for cur != nil {
		nodeAddrA = append(nodeAddrA, cur)
		cur = cur.Next
	}

	cur = headB
	n := len(nodeAddrA)
	for cur != nil {
		for i := 0; i < n; i++ {
			if cur == nodeAddrA[i] {
				return nodeAddrA[i]
			}
		}
		cur = cur.Next
	}

	return nil

}
func main() {
	node5 := &ListNode{5, nil}
	node4 := &ListNode{4, node5}
	node3 := &ListNode{3, node4}
	node2 := &ListNode{2, node3}
	node1 := &ListNode{1, node2}
	node6 :=  &ListNode{6, node4}
	node7 :=  &ListNode{7, node6}

	crossNode := FindCross2(node1, node7)
	if crossNode == nil {
		fmt.Println("cross node nil")
	} else {
		fmt.Printf("crosee node %d\n", crossNode.Val)
	}
}