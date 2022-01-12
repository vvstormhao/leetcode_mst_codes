package main

import (
	"fmt"
	"math"
	"strconv"
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

// 是否是平衡的, 返回二叉树的高度，如果二叉树已经不平衡了则返回-1
// 如果二叉树的左右子树高度之差的绝对值小于1则是平衡的，计算高度，使用后序
func getHeight(cur *TreeNode) int {
	if cur == nil {
		return 0
	}

	heightLeft := getHeight(cur.Left)
	if heightLeft == -1 {
		return -1
	}

	heightRight := getHeight(cur.Right)
	if heightRight == -1 {
		return -1
	}

	if int(math.Abs(float64(heightLeft - heightRight))) > 1 {
		return -1
	}

	return int(math.Max(float64(heightLeft), float64(heightRight))) + 1
}

// 找出二叉树的全部路径
var paths []string
func findAllPath(node *TreeNode, path string) {
	path += strconv.Itoa(node.Val)
	if node.Left == nil && node.Right == nil {
		paths = append(paths, path)
		return
	}

	if node.Left != nil {
		findAllPath(node.Left, path + "->")
	}

	if node.Right != nil {
		findAllPath(node.Right, path + "->")
	}
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

// 合并两个二叉树
func MergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	// 退出条件，一方为空则返回另一方
	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	t1.Val += t2.Val
	t1.Left = MergeTrees(t1.Left, t2.Left)
	t1.Right = MergeTrees(t1.Right, t2.Right)
	return t1
}


// 获取所有左叶子之和
func leftLeafSum(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftSum := leftLeafSum(node.Left)
	rightSum := leftLeafSum(node.Right)

	var midVal int
	if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
		midVal =  node.Left.Val
	}

	return midVal + leftSum + rightSum
}

func searchLevel(node *TreeNode) [][]int{
	if node == nil {
		return nil
	}

	var levelData []int // 缓存每一层的数据
	var allData [][]int // 缓存全部层级的数据
	var deque []*TreeNode
	var popIndex int

	deque = append(deque, node)
	for len(deque) - popIndex > 0 {
		size := len(deque) - popIndex
		for i := 0; i < size; i++ {
			v := deque[popIndex]
			popIndex++

			levelData = append(levelData, v.Val)
			if v.Left != nil {
				deque = append(deque, v.Left)
			}
	
			if v.Right != nil {
				deque = append(deque, v.Right)
			}
		}

		allData = append(allData, levelData)
		levelData = []int{}
	}

	return allData
}
// 找树左下角的值
func getLeftVal(node *TreeNode) int {
	// 使用层序遍历
	levelData := searchLevel(node)
	size := len(levelData)
	return levelData[size-1][0]
}

// 路径总和
func HasPathSum(node *TreeNode, target int) bool {
	if node != nil && target <= 0 {
		return false
	}

	if node == nil {
		if target == 0 {
			return true
		} else {
			return false
		}
	}

	target = target - node.Val
	has := HasPathSum(node.Left, target)
	if has {
		return has
	}

	has = HasPathSum(node.Right, target)
	if has {
		return has
	}
	
	return false
}

// 从中序和后序，构造二叉树，给出前序
func CreateBinary(inOrder []int, postOrder []int) *TreeNode {
	inSize := len(inOrder)
	postSize := len(postOrder)

	if inSize == 0 || postSize == 0 {
		return nil
	}

	rootVal := postOrder[postSize - 1]
	rootNode := &TreeNode{
		Val:rootVal,
	}

	var rootPos int
	for pos, val := range inOrder {
		if val == rootVal {
			rootPos = pos
			break
		}
	}

	leftInOrder := inOrder[:rootPos]
	fmt.Printf("inOrder %v     rootPos %d,  leftInOrder %v\n", inOrder, rootPos, leftInOrder)
	rightInOrder := inOrder[rootPos+1:]
	fmt.Printf("inOrder %v     rootPos %d,  rightInOrder %v\n", inOrder, rootPos, rightInOrder)

	leftPostOrder := postOrder[:rootPos]
	fmt.Printf("postOrder %v     rootPos %d,  leftPostOrder %v\n",postOrder, rootPos, leftPostOrder)
	postLength := inSize - rootPos - 1
	rightPostOrder:= postOrder[rootPos:rootPos + postLength]
	fmt.Printf("postOrder %v     postLength %d,  rightPostOrder %v\n",postOrder, postLength, rightPostOrder)

	rootNode.Left = CreateBinary(leftInOrder, leftPostOrder)
	rootNode.Right = CreateBinary(rightInOrder, rightPostOrder)
	return rootNode
}

// 最大二叉树
func MaxBinary(data []int) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	var maxVal int
	var maxIndex int
	for k, v := range data {
		if v > maxVal {
			maxVal = v
			maxIndex = k
		}
	}

	root := &TreeNode{
		Val : maxVal,
	}

	root.Left = MaxBinary(data[:maxIndex])
	root.Right = MaxBinary(data[maxIndex+1:])

	return root
}

func searchBST(node *TreeNode, target int) *TreeNode {
	if node == nil {
		return nil
	}

	if node.Val == target {
		return node
	}

	if target > node.Val {
		return searchBST(node.Right)
	} else {
		return searchBST(node.Left)
	}
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
/*
	maxDepth := getMaxDepth(tree1)
	fmt.Printf("MaxDepth %d\n", maxDepth)

	minDepth := getMinDepth(tree3)
	fmt.Printf("MinDepth %d\n", minDepth)

	count := getNodeCount(tree2)
	fmt.Printf("node Count %d\n", count)

	newBinary := binaryCombine(tree3, tree3)
	PrintOrderLevel(newBinary)
*/
	findAllPath(tree1, "")
	fmt.Printf("paths %v\n", paths)

	val := getLeftVal(bst1)
	fmt.Printf("left val %v\n", val)

	has := HasPathSum(tree1, 22)
	fmt.Printf("has %v\n", has)

	//inOrder := []int{9,3,15,20,7}
	//postOrder := []int{9,15,7,20,3}

	//tree := CreateBinary(inOrder, postOrder)
	//PrintOrderFirst(tree)

	data := []int{3,2,1,6,0,5}
	tree := MaxBinary(data)
	PrintOrderFirst(tree)
}
