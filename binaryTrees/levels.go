package binaryTrees

type Node struct {
	Load *BinaryTree
	Next *Node
}
type List struct {
	First *Node
	Last  *Node
}

func JoinLevels(bt *BinaryTree) map[int]*List {
	levelLists := make(map[int]*List)
	levelsToLists(bt, 1, levelLists)
	return levelLists
}

func levelsToLists(bt *BinaryTree, level int, acc map[int]*List) {
	if bt == nil {
		return
	}

	levelsToLists(bt.Left, level+1, acc)
	levelsToLists(bt.Right, level+1, acc)

	if _, ok := acc[level]; !ok {
		var seed *Node
		seed = &Node{Load: bt}
		acc[level] = &List{First: seed, Last: seed}
	} else {
		var last *Node
		last = acc[level].Last
		last.Next = &Node{Load: bt}
		acc[level].Last = last.Next
	}

}
