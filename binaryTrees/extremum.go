package binaryTrees

import "errors"

func Min(bt *BinaryTree) (int, error) {
	if bt == nil {
		return 0, errors.New("No Min in the Binary Tree")
	}
	minBt := bt
	for minBt.Left != nil {
		minBt = minBt.Left
	}
	return minBt.Value, nil
}

func Max(bt *BinaryTree) (int, error) {
	if bt == nil {
		return 0, errors.New("No Max in the Binary Tree")
	}
	maxBt := bt
	for maxBt.Right != nil {
		maxBt = maxBt.Right
	}
	return maxBt.Value, nil
}
