package AVLTree

import (
	"fmt"
	"math"
)

/**
 * @Author: Zheng xiaomin
 * @Date: 2020/9/18 21:35 晚上
 */

//AVL tree是为了防止BST退化成链表
// Add      O(logn)
// Contains O(logn)

type AVLTree struct {
	root *node
	size int
}

type node struct {
	value  int
	left   *node
	right  *node
	height int
}

func NewAVLTreeRoot() *AVLTree {
	return &AVLTree{
		root: nil,
		size: 0,
	}
}

func newNode(value int) *node {
	return &node{
		value:  value,
		left:   nil,
		right:  nil,
		height: 1,
	}
}

// 判断该二叉树是不是一颗二分搜索树
func (a *AVLTree) IsBST() bool {
	result := []int{}
	a.inOrder(a.root, result)

	for i := 0; i < len(result); i++ {
		if i+1 != len(result) {
			if result[i+1] < result[i] {
				return false
			}
		}
	}
	return true
}

// 二分搜索树中序遍历
func (a *AVLTree) InOrder(result []int) {
	fmt.Printf("中序遍历:")
	if a.root == nil {
		fmt.Printf(" tree is empty! ")
	} else {
		a.inOrder(a.root, result)
	}
}

func (a *AVLTree) inOrder(node *node, result []int) {
	if node == nil {
		return
	}

	a.inOrder(node.left, result)
	result = append(result, node.value)
	//fmt.Printf("%d", node.value)
	a.inOrder(node.right, result)
}

// 判断该二叉树是不是一颗平衡二叉树
func (a *AVLTree) IsBalancedTree() bool {
	return a.isBalancedTree(a.root)
}

func (a *AVLTree) isBalancedTree(node *node) bool {
	if node == nil {
		return true
	}

	balancedFactor := node.getBalanceFactor()
	if math.Abs(float64(balancedFactor)) > 1 {
		return false
	}

	return a.isBalancedTree(node.left) && a.isBalancedTree(node.right)
}

func (a *AVLTree) Add(value int) {
	a.root = a.add(value, a.root)
}

func (a *AVLTree) add(value int, node *node) *node {
	if node == nil {
		return newNode(value)
	}

	if node.value < value {
		node.right = a.add(value, node.right)
	} else {
		node.left = a.add(value, node.left)
	}

	// 更新高度
	node.height = 1 + max(node.right.getHeight(), node.left.getHeight())

	// 计算平衡因子
	balanceFactor := node.getBalanceFactor()
	//if math.Abs(float64(balanceFactor)) > 1 {
	//	fmt.Println("unbalanced node :", node.value, " balanceFactor :", balanceFactor)
	//}

	// 维护平衡性
	// LL
	// 对节点y进行向右旋转操作，返回旋转后新的根节点
	//           y                   x
	//          / \                /   \
	//         x  T4    ---->     z     y
	//        / \                / \   / \
	//       z  T3             T1  T2 T3 T4
	//      / \
	//     T1 T2
	if balanceFactor > 1 && node.left.getBalanceFactor() >= 0 {
		return a.rightRotate(node)
	}

	// RR
	// 对节点y进行向左旋转操作，返回旋转后新的根节点
	//           y                       x
	//          / \                    /   \
	//         T1  x    ---->         y     z
	//        	  / \                / \   / \
	//           T2  z             T1  T2 T3 T4
	//              / \
	//             T3 T4
	if balanceFactor < -1 && node.right.getBalanceFactor() <= 0 {
		return a.leftRotate(node)
	}

	// LR
	// 对节点x进行向左旋转操作，再对z进行右旋转
	//           y                   z
	//          / \                /   \
	//         x  T4    ---->     x     y
	//        / \                / \   / \
	//       T1  z             T1  T2 T3 T4
	//      	/ \
	//     	   T2 T3
	if balanceFactor > 1 && node.left.getBalanceFactor() < 0 {
		node.left = a.leftRotate(node.left)
		return a.rightRotate(node)
	}

	// RL
	// 对节点x进行向右旋转操作，再对z进行左旋转
	//           y                     z
	//          / \                  /   \
	//         T1  x    ---->       y     x
	//        	  / \              / \   / \
	//           z  T4           T1  T2 T3 T4
	//         / \
	//        T2 T3
	if balanceFactor < -1 && node.left.getBalanceFactor() > 0 {
		node.right = a.rightRotate(node.right)
		return a.leftRotate(node)
	}

	return node
}

// LL
func (a *AVLTree) rightRotate(y *node) *node {
	x := y.left
	T3 := x.right

	x.right = y
	y.left = T3

	// 更新height
	y.height = max(y.left.getHeight(), y.right.getHeight()) + 1
	x.height = max(x.left.getHeight(), x.right.getHeight()) + 1

	return x
}

// RR
func (a *AVLTree) leftRotate(y *node) *node {
	x := y.right
	T2 := x.left

	x.left = y
	y.right = T2

	// 更新height
	y.height = max(y.left.getHeight(), y.right.getHeight()) + 1
	x.height = max(x.left.getHeight(), x.right.getHeight()) + 1

	return x
}

func (n *node) getBalanceFactor() int {
	if n == nil {
		return 0
	}

	return n.left.getHeight() - n.right.getHeight()
}

func (a *AVLTree) GetSize() int {
	return a.size
}

func (a *AVLTree) IsEmpty() bool {
	return a.size == 0
}

// 获得节点的高度值
func (n *node) getHeight() int {
	if n == nil {
		return 0
	}

	return n.height
}

func (a *AVLTree) Contains(value int) bool {
	return a.contains(value, a.root)
}

func (a *AVLTree) contains(value int, node *node) bool {
	if node == nil {
		return false
	}

	if node.value != value {
		if node.value > value {
			return a.contains(value, node.left)
		} else {
			return a.contains(value, node.right)
		}
	} else {
		return true
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func (a *AVLTree) Remove(value int) {
	node := a.GetNode(value)
	if node == nil {
		return
	}

	a.root = a.remove(a.root, value)
	return
}

func (a *AVLTree) remove(n *node, value int) *node {
	if n == nil {
		return nil
	}

	var tmpNode *node

	if n.value > value {
		n.right = a.remove(n.right, value)
		tmpNode = n
	} else if n.value < value {
		n.left = a.remove(n.left, value)
		tmpNode = n
	} else {

		// 左子树为空
		if n.left == nil {
			rightNode := n.right
			n.right = nil
			a.size--
			tmpNode = rightNode
		}

		// 右子树为空
		if n.right == nil {
			leftNode := n.left
			n.left = nil
			a.size--
			tmpNode = leftNode
		}

		successor := a.minimum(n.right)
		successor.right = a.remove(n.right, value)
		successor.left = n.left
		n.left, n.right = nil, nil

		tmpNode = successor

	}

	if tmpNode == nil {
		return nil
	}

	// 更新高度
	tmpNode.height = 1 + max(tmpNode.right.getHeight(), tmpNode.left.getHeight())

	// 计算平衡因子
	balanceFactor := tmpNode.getBalanceFactor()

	// 维护平衡性
	// LL
	if balanceFactor > 1 && tmpNode.left.getBalanceFactor() >= 0 {
		return a.rightRotate(tmpNode)
	}

	// RR
	if balanceFactor < -1 && tmpNode.right.getBalanceFactor() <= 0 {
		return a.leftRotate(tmpNode)
	}

	// LR
	if balanceFactor > 1 && tmpNode.left.getBalanceFactor() < 0 {
		tmpNode.left = a.leftRotate(tmpNode.left)
		return a.rightRotate(tmpNode)
	}

	// RL
	if balanceFactor < -1 && tmpNode.left.getBalanceFactor() > 0 {
		tmpNode.right = a.rightRotate(tmpNode.right)
		return a.leftRotate(tmpNode)
	}

	return tmpNode
}

func (a *AVLTree) GetNode(value int) *node {
	if a.root == nil {
		return nil
	}

	return a.getNode(a.root, value)
}

func (a *AVLTree) getNode(node *node, value int) *node {
	if node == nil {
		return nil
	}

	if node.value < value {
		return a.getNode(node.right, value)
	} else if node.value > value {
		return a.getNode(node.left, value)
	}

	return node
}

func (a *AVLTree) FindMinimum() (*node, error) {
	if a.root == nil {
		return nil, avlIsEmpty
	}

	return a.minimum(a.root), nil
}

func (a *AVLTree) minimum(node *node) *node {
	for ; node != nil; {
		if node.left == nil {
			break
		}
		node = node.left
	}
	return node
}

func (a *AVLTree) removeMin(removeNode *node) *node {
	if removeNode.left == nil {
		rightChild := removeNode.right
		removeNode.right = nil
		a.size--
		return rightChild
	}

	removeNode.left = a.removeMin(removeNode.left)
	return removeNode
}
