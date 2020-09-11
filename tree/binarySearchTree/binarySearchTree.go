package binarySearchTree

import "fmt"

type BinarySearchTree struct {
	root *node
	size int
}

type node struct {
	value int
	left  *node
	right *node
}

func NewBinarySearchTreeRoot() *BinarySearchTree {
	return &BinarySearchTree{
		root: nil,
		size: 0,
	}
}

func newNode(value int) *node {
	return &node{
		value: value,
		left:  nil,
		right: nil,
	}
}

func (b *BinarySearchTree) Add(value int) {
	b.root = b.add(value, b.root)
}

func (b *BinarySearchTree) add(value int, node *node) *node {
	if node == nil {
		return newNode(value)
	}

	if node.value < value {
		node.right = b.add(value, node.right)
	} else {
		node.left = b.add(value, node.left)
	}

	return node
}

func (b *BinarySearchTree) Contains(value int) bool {
	return b.contains(value, b.root)
}

func (b *BinarySearchTree) contains(value int, node *node) bool {
	if node == nil {
		return false
	}

	if node.value != value {
		if node.value > value {
			return b.contains(value, node.left)
		} else {
			return b.contains(value, node.right)
		}
	} else {
		return true
	}
}

// 二分搜索树前序遍历
func (b *BinarySearchTree) PreOrder() {
	fmt.Printf("前序遍历:")
	if b.root == nil {
		fmt.Printf(" tree is empty! ")
	} else {
		b.preOrder(b.root)
	}

	fmt.Println()
}

func (b *BinarySearchTree) preOrder(node *node) {
	if node == nil {
		return
	}

	fmt.Printf("%d", node.value)
	b.preOrder(node.left)
	b.preOrder(node.right)
}

// 二分搜索树后序遍历
func (b *BinarySearchTree) PostOrder() {
	fmt.Printf("后序遍历:")
	if b.root == nil {
		fmt.Printf(" tree is empty! ")
	} else {
		b.postOrder(b.root)
	}

	fmt.Println()
}

func (b *BinarySearchTree) postOrder(node *node) {
	if node == nil {
		return
	}

	b.postOrder(node.left)
	b.postOrder(node.right)
	fmt.Printf("%d", node.value)
}

// 二分搜索树中序遍历
func (b *BinarySearchTree) InOrder() {
	fmt.Printf("中序遍历:")
	if b.root == nil {
		fmt.Printf(" tree is empty! ")
	} else {
		b.inOrder(b.root)
	}

	fmt.Println()

}

func (b *BinarySearchTree) inOrder(node *node) {
	if node == nil {
		return
	}

	b.inOrder(node.left)
	fmt.Printf("%d", node.value)
	b.inOrder(node.right)
}
