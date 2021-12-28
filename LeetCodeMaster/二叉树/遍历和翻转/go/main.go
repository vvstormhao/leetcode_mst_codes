package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 前序遍历——递归方法
func PrintOrderFirst(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Printf("%d ", root.Val)
	PrintOrderFirst(root.Left)
	PrintOrderFirst(root.Right)
}

// 中序遍历——递归方法
func PrintOrderMid(root *TreeNode) {
	if root == nil {
		return
	}

	PrintOrderMid(root.Left)
	fmt.Printf("%d ", root.Val)
	PrintOrderMid(root.Right)
}

// 后序遍历——递归方法
func PrintOrderLast(root *TreeNode) {
	if root == nil {
		return
	}

	PrintOrderLast(root.Left)
	PrintOrderLast(root.Right)
	fmt.Printf("%d ", root.Val)
}

// 层级遍历——迭代方法
// 使用队列思路
func PrintOrderLevel(root *TreeNode){
	if root == nil {
		return
	}

	var deque []*TreeNode
	indexPop := 0
	var levels[][]interface{}
	var level []interface{}
	deque = append(deque, root)
	level = append(level, root.Val)
	levels = append(levels, level)
	for len(deque) - indexPop > 0 {
		level = []interface{}{}
		size := len(deque) - indexPop
		for i := 0; i <  size; i++{
			node := deque[indexPop]
			indexPop++
			if node.Left != nil {
				deque = append(deque, node.Left)
				level = append(level, node.Left.Val)
			} else {
				level = append(level, nil)
			}
	
			if node.Right != nil {
				deque = append(deque, node.Right)
				level = append(level, node.Right.Val)
			} else {
				level = append(level, nil)
			}
		}

		levels = append(levels, level)
	}

	fmt.Printf("%v\n",levels)
}

// 翻转二叉树
func Revert(root *TreeNode) {
	if root == nil {
		return
	}

	Revert(root.Left)
	Revert(root.Right)
	tmp := root.Right
	root.Right = root.Left
	root.Left = tmp
}

// 判断二叉树是否对称
func IsMirror(left *TreeNode, right *TreeNode) bool {
	if left == nil && right != nil {
		return false
	}

	if left != nil && right == nil {
		return false
	}

	if left == nil && right == nil {
		return true
	}

	if left.Val != right.Val {
		return false
	}

	outside := IsMirror(left.Left, right.Right)
	inside := IsMirror(left.Right, right.Left)
	return outside && inside
}

// 获取二叉树的最大深度
func getMaxDepth(root *TreeNode) int{
	if root == nil {
		return 0
	}

	depLeft := getMaxDepth(root.Left)
	depRight := getMaxDepth(root.Right)

	var max int
	if depLeft >= depRight {
		max = depLeft
	} else {
		max = depRight
	}

	return 1 + max
}

// 获取二叉树的最小深度
func getMinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depLeft := getMinDepth(root.Left)
	depRight := getMinDepth(root.Right)

	var min int
	if depLeft <= depRight {
		min = depLeft
	} else {
		min = depRight
	}

	return 1 + min
}

// 获取节点总数
func getNodeCount(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 +  getNodeCount(root.Left) + getNodeCount(root.Right)
}

func main() {

	tree2 := &TreeNode{
		Val:0,
		Left:&TreeNode{
			Val:1,
			Left:&TreeNode{
				Val:3,
				Left:&TreeNode{
					Val:7,
				},
				Right:&TreeNode{
					Val:8,
				},
			},
			Right:&TreeNode{
				Val:4,
				Left:&TreeNode{
					Val:9,
				},
				Right:&TreeNode{
					Val:10,
				},
			},
		},
		Right:&TreeNode{
			Val:2,
			Left:&TreeNode{
				Val:5,
			},
			Right:&TreeNode{
				Val:6,
			},
		},
	}

	tree1 := &TreeNode{
		Val:0,
		Left:&TreeNode{
			Val:2,
			Left:&TreeNode{
				Val:3,
				Left:&TreeNode{
					Val:7,
				},
				Right:&TreeNode{
					Val:8,
				},
			},
			Right:&TreeNode{
				Val:4,
				Left:&TreeNode{
					Val:9,
				},
				Right:&TreeNode{
					Val:10,
					Left:&TreeNode{
						Val:11,
						Right:&TreeNode{
							Val:12,
						},
					},
				},
			},
		},
		Right:&TreeNode{
			Val:2,
			Left:&TreeNode{
				Val:4,
				Left:&TreeNode{
					Val:10,
				},
				Right:&TreeNode{
					Val:9,
				},
			},
			Right:&TreeNode{
				Val:3,
				Left:&TreeNode{
					Val:8,
				},
				Right:&TreeNode{
					Val:7,
				},
			},
		},
	}

	/*
		    0
	    1        2
    3      4	5   6
  7    8 9   10   
          11  
		      12
	
	*/
	/*
	       0
		2        1
	6      5  4      3
	        9   10 8    7
	*/

	/*
	PrintOrderFirst(node0)
	fmt.Printf("\n")
	PrintOrderMid(node0)
	fmt.Printf("\n")
	PrintOrderLast(node0)
	fmt.Printf("\n")
	*/
	//PrintOrderLevel(tree1)
	//fmt.Printf("\n")

	//Revert(treeMirro)
	//PrintOrderLevel(treeMirro)
	//fmt.Printf("\n")
	
	//isMir := IsMirror(tree1, treeMirro)
	//fmt.Printf("isMirr %v\n",isMir)

	maxDepth := getMaxDepth(tree1)
	fmt.Printf("MaxDepth %d\n", maxDepth)

	tree3 := &TreeNode{
		Val:0,
		Right:&TreeNode{
			Val:1,
			Right:&TreeNode{
				Val:2,
				Right:&TreeNode{
					Val:3,
				},
			},
			Left:&TreeNode{
				Val:4,
			},
		},
		Left:&TreeNode{
			Val:3,
		},
	}
	minDepth := getMinDepth(tree3)
	fmt.Printf("MinDepth %d\n", minDepth)

	count := getNodeCount(tree2)
	fmt.Printf("node Count %d\n", count)

}