package binarySearchTree

import (
	"fmt"
	"testing"
)

func TestNode_Add(t *testing.T) {
	bst := NewBinarySearchTreeRoot()

	//a := []int{6, 5, 2, 1, 3, 4, 8, 7, 10, 9}
	a := []int{5, 3, 2, 4, 6, 8}
	//////////////////
	//       5      //
	//     /  \     //
	//    3    6    //
	//   /\     \   //
	//  2  4     8  //
	//////////////////
	for i := 0; i < len(a); i++ {
		bst.Add(a[i])
	}

	fmt.Printf("contains : %v \n", bst.Contains(-1))

	//bst.PostOrder()
	//bst.PreOrder()
	//bst.PreOrderNR()
	//bst.InOrder()
	//bst.InOrderNR()

	//bst.LevelOrder()

	//maxNode, _ := bst.FindMaximum()
	//t.Logf("max : %d", maxNode.value)
	//minNode, _ := bst.FindMinimum()
	//t.Logf("min : %d", minNode.value)

	//_, _ = bst.RemoveMin()
	//_, _ = bst.RemoveMin()

	_ = bst.Remove(3)

	bst.LevelOrder()
}
