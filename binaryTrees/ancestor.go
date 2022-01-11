package binaryTrees

func CommonAncestor(bt1 *BinaryTree2, bt2 *BinaryTree2) *BinaryTree2 {
	if bt1 == nil || bt2 == nil {
		return nil
	}
	var p1 *BinaryTree2
	var p2 *BinaryTree2
	p1 = bt1
	p2 = bt2
	//fmt.Printf("> bt1=[%d], bt2=[%d]\n", bt1.Value, bt2.Value)
	if bt1.Value > bt2.Value {
		p1 = bt2
		p2 = bt1
	}
	//fmt.Printf("> p1=[%d], p2=[%d]\n", p1.Value, p2.Value)
	var q *BinaryTree2
	var prev *BinaryTree2
	var prevAcc *BinaryTree2
	q = p1.Parent
	prev = p1
	prevAcc = p1
	//fmt.Printf("> q=[%d], prev=[%d]\n", q.Value, prev.Value)
	if q != nil {
		//fmt.Println("in if")
		for q != nil && q.Value <= p2.Value {
			if q.Left == prev {
				prevAcc = q
			}
			prev = q
			q = q.Parent
			//fmt.Printf("q=[%v], prev=[%v], prevAcc=[%v]\n", q, prev.Value, prevAcc.Value)
		}
	}
	if Find2(prevAcc, p2) {
		//fmt.Println("return-1")
		return prevAcc
	} else if Find2(p1, p2) {
		//fmt.Println("return-2")
		return p1
	} else {
		//fmt.Println("return-3")
		return nil
	}
}
