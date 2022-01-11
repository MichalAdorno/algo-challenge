package binaryTrees

import "testing"

func Test_CommonAncestor(t *testing.T) {
	//prepare
	var bt *BinaryTree2
	bt = b47
	bt.Right = b50
	bt.Left = b30
	b30.Left = b20
	b30.Right = b45
	b45.Left = b37
	b45.Right = b46
	b37.Left = b35
	b37.Right = b39
	b39.Left = b38
	b35.Left = b31
	b35.Right = b36
	b31.Right = b34
	b20.Right = b29

	b50.Parent = b47
	b30.Parent = b47
	b45.Parent = b30
	b20.Parent = b30
	b29.Parent = b20
	b37.Parent = b45
	b46.Parent = b45
	b35.Parent = b37
	b39.Parent = b37
	b38.Parent = b39
	b31.Parent = b35
	b36.Parent = b35
	b34.Parent = b31
	//test
	var tests = []struct {
		name     string
		p1       *BinaryTree2
		p2       *BinaryTree2
		ancestor *BinaryTree2
	}{
		{
			name:     "34-38->37",
			p1:       b34,
			p2:       b38,
			ancestor: b37,
		},
		{
			name:     "34-31->31",
			p1:       b34,
			p2:       b31,
			ancestor: b31,
		},
		{
			name:     "34-50->47",
			p1:       b34,
			p2:       b50,
			ancestor: b47,
		},
		{
			name:     "20-29->20",
			p1:       b20,
			p2:       b29,
			ancestor: b20,
		},
		{
			name:     "31-36->35",
			p1:       b31,
			p2:       b36,
			ancestor: b35,
		},
		{
			name:     "29-36->30",
			p1:       b29,
			p2:       b36,
			ancestor: b30,
		},
		{
			name:     "29-50->47",
			p1:       b29,
			p2:       b50,
			ancestor: b47,
		},
		{
			name:     "29-30->30",
			p1:       b29,
			p2:       b30,
			ancestor: b30,
		},
		{
			name:     "34-50->47",
			p1:       b34,
			p2:       b50,
			ancestor: b47,
		},
		{
			name:     "39-50->47",
			p1:       b39,
			p2:       b50,
			ancestor: b47,
		},
		{
			name:     "36-46->45",
			p1:       b36,
			p2:       b46,
			ancestor: b45,
		},
		{
			name:     "34-50->47",
			p1:       b34,
			p2:       b50,
			ancestor: b47,
		},
		{
			name:     "31-45->45",
			p1:       b31,
			p2:       b45,
			ancestor: b45,
		},
		{
			name:     "38-47->47",
			p1:       b38,
			p2:       b47,
			ancestor: b47,
		},
		{
			name:     "36-38->37",
			p1:       b36,
			p2:       b38,
			ancestor: b37,
		},
	}
	for _, test := range tests {
		if res := CommonAncestor(test.p1, test.p2); res == nil || res.Value != test.ancestor.Value {
			if res == nil {
				t.Errorf("The CommonAncestor of [%d,%d] is: expected [%d] but got NIL ",
					test.p1.Value, test.p2.Value, test.ancestor.Value)
			} else {
				t.Errorf("The CommonAncestor of [%d,%d] is: expected [%d] but got [%d] ",
					test.p1.Value, test.p2.Value, test.ancestor.Value, res.Value)
			}

		}
	}
}

var (
	b60 = NewBinaryTree2(60)
	b50 = NewBinaryTree2(50)
	b47 = NewBinaryTree2(47)
	b46 = NewBinaryTree2(46)
	b45 = NewBinaryTree2(45)
	b40 = NewBinaryTree2(40)
	b39 = NewBinaryTree2(39)
	b38 = NewBinaryTree2(38)
	b37 = NewBinaryTree2(37)
	b36 = NewBinaryTree2(36)
	b35 = NewBinaryTree2(35)
	b34 = NewBinaryTree2(34)
	b31 = NewBinaryTree2(31)
	b30 = NewBinaryTree2(30)
	b29 = NewBinaryTree2(29)
	b20 = NewBinaryTree2(20)
)
