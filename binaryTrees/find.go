package binaryTrees

func Find(bt *BinaryTree, value int) bool {
	if bt == nil {
		return false
	}
	btUnderSearch := bt
	for btUnderSearch != nil {
		if value < btUnderSearch.Value {
			btUnderSearch = btUnderSearch.Left
		} else if btUnderSearch.Value < value {
			btUnderSearch = btUnderSearch.Right
		} else {
			return true
		}
	}
	return btUnderSearch != nil && btUnderSearch.Value == value
}
func Find2(bt *BinaryTree2, target *BinaryTree2) bool {
	if bt == nil {
		return false
	}
	btUnderSearch := bt
	for btUnderSearch != nil {
		if target.Value < btUnderSearch.Value {
			btUnderSearch = btUnderSearch.Left
		} else if btUnderSearch.Value < target.Value {
			btUnderSearch = btUnderSearch.Right
		} else {
			return true
		}
	}
	return btUnderSearch != nil && btUnderSearch.Value == target.Value
}
