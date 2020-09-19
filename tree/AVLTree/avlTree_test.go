package AVLTree

import "testing"

func TestAVLTree(t *testing.T) {
	avlTree := NewAVLTreeRoot()
	tmpArray := []int{}

	for i := 0; i < 10; i++ {
		avlTree.Add(i)
		tmpArray = append(tmpArray, i)
	}

	t.Logf("is BST tree : %v", avlTree.IsBST())
	t.Logf("is AVL tree : %v", avlTree.IsBalancedTree())
	t.Logf("height : %d", avlTree.root.getHeight())
}

func BenchmarkAVLTree_Add(b *testing.B) {
	avlTree := NewAVLTreeRoot()
	tmpArray := []int{}

	times := 100000

	for i := 0; i < times; i++ {
		avlTree.Add(i)
		tmpArray = append(tmpArray, i)
	}
}

func BenchmarkAVLTree_Contains(b *testing.B) {
	avlTree := NewAVLTreeRoot()
	tmpArray := []int{}

	times := 100000

	for i := 0; i < times; i++ {
		avlTree.Add(i)
		tmpArray = append(tmpArray, i)
	}

	b.ResetTimer()

	for i := 0; i < times; i++ {
		_ = avlTree.Contains(i)
	}
}
