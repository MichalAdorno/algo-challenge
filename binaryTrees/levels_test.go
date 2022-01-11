package binaryTrees

import (
	"fmt"
	"testing"
)

func Test_JoinLevels(t *testing.T) {
	//prepare
	var lbt *BinaryTree
	lbt = lb47
	lbt.Right = lb50
	lbt.Left = lb30
	lb50.Right = lb70
	lb30.Left = lb20
	lb30.Right = lb45
	lb45.Left = lb37
	lb45.Right = lb46
	lb37.Left = lb35
	lb37.Right = lb39
	lb39.Left = lb38
	lb35.Left = lb31
	lb35.Right = lb36
	lb31.Right = lb34
	lb20.Right = lb29

	//test
	var tests = []struct {
		name   string
		p1     *BinaryTree
		height int
	}{
		{
			name:   "example",
			p1:     lbt,
			height: 7,
		},
	}
	for _, test := range tests {
		if res := JoinLevels(test.p1); res != nil {
			for i := 1; i <= test.height; i++ {
				fmt.Printf("level = [%d]: ", i)
				var iter *Node = res[i].First
				for iter != nil {
					fmt.Printf("[%d]", iter.Load.Value)
					iter = iter.Next
				}
				fmt.Println("")
			}
		}
	}
}

func NewBinaryTree(value int) *BinaryTree {
	return &BinaryTree{Value: value}
}

var (
	lb70 = NewBinaryTree(70)
	lb60 = NewBinaryTree(60)
	lb50 = NewBinaryTree(50)
	lb47 = NewBinaryTree(47)
	lb46 = NewBinaryTree(46)
	lb45 = NewBinaryTree(45)
	lb40 = NewBinaryTree(40)
	lb39 = NewBinaryTree(39)
	lb38 = NewBinaryTree(38)
	lb37 = NewBinaryTree(37)
	lb36 = NewBinaryTree(36)
	lb35 = NewBinaryTree(35)
	lb34 = NewBinaryTree(34)
	lb31 = NewBinaryTree(31)
	lb30 = NewBinaryTree(30)
	lb29 = NewBinaryTree(29)
	lb20 = NewBinaryTree(20)
)
