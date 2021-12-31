package main

import (
	"fmt"
)

/*
tree1 完全二叉树 非二叉查找数
             1
	     /       \
	    4         2
	  /    \     / \
	 3      6   13  11
    / \    / \
   7   8  9   10
*/
var tree1 *TreeNode = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 10,
			},
		},
	},
	Right: &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 13,
		},
		Right: &TreeNode{
			Val: 11,
		},
	},
}

/*
bst1 平衡二叉查找树
         10
		/   \
	   3    19
	  / \   / \
	 1   6 11  20
*/

var bst1 *TreeNode = &TreeNode{
	Val: 10,
	Left: &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 6,
		},
	},
	Right: &TreeNode{
		Val: 19,
		Left: &TreeNode{
			Val: 11,
		},
		Right: &TreeNode{
			Val: 20,
		},
	},
}

var tree2 *TreeNode = &TreeNode{
	Val: 0,
	Left: &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 10,
			},
		},
	},
	Right: &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 5,
		},
		Right: &TreeNode{
			Val: 6,
		},
	},
}

/*
    5
   / \
  3   1
     / \
	4   2
	     \
		   3
*/
var tree3 *TreeNode = &TreeNode{
	Val: 5,
	Left: &TreeNode{
		Val: 3,
	},
	Right: &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
	},
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
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
func PrintOrderLevel(root *TreeNode) {
	if root == nil {
		return
	}

	var deque []*TreeNode
	indexPop := 0
	var levels [][]interface{}
	var level []interface{}
	deque = append(deque, root)
	level = append(level, root.Val)
	levels = append(levels, level)
	for len(deque)-indexPop > 0 {
		level = []interface{}{}
		size := len(deque) - indexPop
		for i := 0; i < size; i++ {
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

	fmt.Printf("%v\n", levels)
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
func getMaxDepth(root *TreeNode) int {
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

	return 1 + getNodeCount(root.Left) + getNodeCount(root.Right)
}

// 合并两个二叉树,合并后二叉树的节点为两个节点数值之和
func binaryCombine(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}

	var node *TreeNode
	if root1 != nil && root2 != nil {
		node = &TreeNode{
			Val: root1.Val + root2.Val,
		}
		node.Left = binaryCombine(root1.Left, root2.Left)
		node.Right = binaryCombine(root1.Right, root2.Right)
	} else if root1 != nil && root2 == nil {
		node = &TreeNode{
			Val: root1.Val,
		}
		node.Left = binaryCombine(root1.Left, nil)
		node.Right = binaryCombine(root1.Right, nil)
	} else {
		node = &TreeNode{
			Val: root2.Val,
		}
		node.Left = binaryCombine(nil, root2.Left)
		node.Right = binaryCombine(nil, root2.Right)
	}

	return node
}

// 判断是否为二叉搜索树
// 二叉搜索树的最小绝对值差
// 获取众数
// 众数即出现频率最多的节点集合
// 获取公共祖先

func main() {

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

	minDepth := getMinDepth(tree3)
	fmt.Printf("MinDepth %d\n", minDepth)

	count := getNodeCount(tree2)
	fmt.Printf("node Count %d\n", count)

	newBinary := binaryCombine(tree3, tree3)
	PrintOrderLevel(newBinary)
}
