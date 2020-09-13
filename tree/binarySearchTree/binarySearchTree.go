package binarySearchTree

import (
	"fmt"
	"github.com/navi-tt/data-structure/queue"
	"github.com/navi-tt/data-structure/stack"
)

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

// 二分搜索树前序遍历非递归实现
func (b *BinarySearchTree) PreOrderNR() {
	fmt.Printf("非递归前序遍历:")
	if b.root == nil {
		fmt.Printf(" tree is empty! ")
		return
	}

	linkedListStack := stack.NewLinkedListStack()

	_ = linkedListStack.Push(b.root)

	for ; !linkedListStack.IsEmpty(); {
		v, _ := linkedListStack.Pop()

		nv := v.(*node)

		fmt.Printf("%d", nv.value)

		if nv.right != nil {
			_ = linkedListStack.Push(nv.right)
		}

		if nv.left != nil {
			_ = linkedListStack.Push(nv.left)
		}
	}

	fmt.Println("")
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

//二分搜索树中序遍历非递归实现
func (b *BinarySearchTree) InOrderNR() {
	fmt.Printf("非递归中序遍历:")
	if b.root == nil {
		fmt.Printf(" tree is empty! ")
		return
	}

	linkedListStack := stack.NewLinkedListStack()
	_ = linkedListStack.Push(b.root)
	leftChild := b.root.left

	for ; !linkedListStack.IsEmpty(); {

		for ; leftChild != nil; {
			_ = linkedListStack.Push(leftChild)
			leftChild = leftChild.left
		}

		v, _ := linkedListStack.Pop()

		nv := v.(*node)

		fmt.Printf("%d", nv.value)

		if nv.right != nil {
			_ = linkedListStack.Push(nv.right)
		}
	}

	fmt.Println("")
}

//二分搜索树层级遍历
func (b *BinarySearchTree) LevelOrder() {
	fmt.Printf("层序遍历:")
	if b.root == nil {
		fmt.Printf(" tree is empty! ")
		return
	}

	linkedListQueue := queue.NewLinkedListQueue()
	_ = linkedListQueue.Enqueue(b.root)
	for ; !linkedListQueue.IsEmpty(); {

		v, _ := linkedListQueue.Dequeue()
		nv := v.(*node)

		fmt.Printf("%d", nv.value)

		if nv.left != nil {
			_ = linkedListQueue.Enqueue(nv.left)
		}

		if nv.right != nil {
			_ = linkedListQueue.Enqueue(nv.right)
		}

	}

	fmt.Println()
}

//找到二分搜索树的最小值
func (b *BinarySearchTree) FindMinimum() (*node, error) {
	if b.root == nil {
		return nil, bstIsEmpty
	}

	return b.minimum(b.root), nil
}

func (b *BinarySearchTree) minimum(node *node) *node {
	for ; node != nil; {
		if node.left == nil {
			break
		}
		node = node.left
	}
	return node
}

//找到二分搜索树的最大值
func (b *BinarySearchTree) FindMaximum() (*node, error) {
	if b.root == nil {
		return nil, bstIsEmpty
	}

	return b.maximum(b.root), nil
}

func (b *BinarySearchTree) maximum(node *node) *node {
	for ; node != nil; {
		if node.right == nil {
			break
		}
		node = node.right
	}
	return node
}

//删除二分搜索树种最小值，并返回删除后的根节点
func (b *BinarySearchTree) RemoveMin() (*node, error) {
	minNode, err := b.FindMinimum()
	if err != nil {
		return nil, err
	}

	b.root = b.removeMin(b.root)
	return minNode, nil
}

func (b *BinarySearchTree) removeMin(removeNode *node) *node {
	if removeNode.left == nil {
		rightChild := removeNode.right
		removeNode.right = nil
		b.size--
		return rightChild
	}

	removeNode.left = b.removeMin(removeNode.left)
	return removeNode
}

//删除二分搜索树种最大值，并返回删除后的根节点
func (b *BinarySearchTree) RemoveMax() (*node, error) {
	maxNode, err := b.FindMaximum()
	if err != nil {
		return nil, err
	}

	b.root = b.removeMax(b.root)
	return maxNode, nil
}

func (b *BinarySearchTree) removeMax(removeNode *node) *node {
	if removeNode.left == nil {
		leftChild := removeNode.left
		removeNode.left = nil
		b.size--
		return leftChild
	}

	removeNode.right = b.removeMax(removeNode.right)
	return removeNode
}

func (b *BinarySearchTree) Remove(value int) error {
	if b.root == nil {
		return bstIsEmpty
	}

	b.root = b.remove(b.root, value)
	return nil
}

//删除以node为根的值为value的节点，返回删除节点后新的二分搜索树的跟
func (b *BinarySearchTree) remove(node *node, value int) *node {
	if node == nil {
		return nil
	}

	if node.value > value {
		node.left = b.remove(node.left, value)
		return node
	} else if node.value < value {
		node.right = b.remove(node.right, value)
		return node
	} else {
		//待删除节点左子树为空
		if node.left == nil {
			rightChild := node.right
			node.left = nil
			b.size--
			return rightChild
		}

		//待删除节点右子树为空
		if node.right == nil {
			leftChild := node.left
			node.right = nil
			b.size--
			return leftChild
		}

		//待删除节点左右子树都不为空
		//找到删除节点右子树最小值的节点
		//用这个节点顶替要被删除的节点
		rightMinNode := b.minimum(node.right)
		rightMinNode.right = b.removeMin(node.right)
		rightMinNode.left = node.left

		node.right = nil
		node.left = nil

		return rightMinNode
	}

}
