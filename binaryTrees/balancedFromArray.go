package binaryTrees

import "fmt"

func BuildMinimalTree(arr []int) *BinaryTree {
	return createMinimalBST(arr, 0, len(arr)-1)
}

func createMinimalBST(arr []int, start int, end int) *BinaryTree {
	if start > end {
		return nil
	}
	mid := (start + end) / 2
	fmt.Printf("start=[%v], end=[%v], mid=[%v], arr[mid]=[%v]\n", start, end, mid, arr[mid])
	bt := &BinaryTree{
		Value: arr[mid],
		Left:  createMinimalBST(arr, start, mid-1),
		Right: createMinimalBST(arr, mid+1, end),
	}
	return bt
}

func InOrder(bt *BinaryTree) []int {
	acc := make([]int, 0)
	inOrder(bt, &acc)
	fmt.Printf("++++[%v]\n", acc)
	return acc
}
func inOrder(bt *BinaryTree, acc *[]int) {
	if bt == nil {
		return
	}
	inOrder(bt.Left, acc)
	fmt.Printf("bt.Value=[%v], acc=[%v]\n", bt.Value, acc)
	*acc = append(*acc, bt.Value)
	inOrder(bt.Right, acc)

}
