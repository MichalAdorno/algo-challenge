package binaryTrees

func Next(bt *BinaryTree) *BinaryTree {
	if bt == nil || bt.Right == nil {
		return nil
	}
	nextBt := bt.Right
	for nextBt.Left != nil {
		nextBt = nextBt.Left
	}
	return nextBt
}
