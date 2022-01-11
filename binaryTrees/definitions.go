package binaryTrees

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}
type BinaryTree2 struct {
	Value  int
	Left   *BinaryTree2
	Right  *BinaryTree2
	Parent *BinaryTree2
}

func NewBinaryTree2(value int) *BinaryTree2 {
	return &BinaryTree2{Value: value}
}

var (
	bt2 = BinaryTree{
		Left: &BinaryTree{
			Left: &BinaryTree{
				Left:  nil,
				Value: 20,
				Right: &BinaryTree{
					Left:  nil,
					Value: 29,
					Right: nil,
				},
			},
			Value: 30,
			Right: &BinaryTree{
				Left: &BinaryTree{
					Left: &BinaryTree{
						Left: &BinaryTree{
							Left:  nil,
							Value: 31,
							Right: &BinaryTree{
								Left:  nil,
								Value: 34,
								Right: nil,
							},
						},
						Value: 35,
						Right: &BinaryTree{
							Left:  nil,
							Value: 36,
							Right: nil,
						},
					},
					Value: 37,
					Right: &BinaryTree{
						Left: &BinaryTree{
							Left:  nil,
							Value: 38,
							Right: nil,
						},
						Value: 39,
						Right: nil,
					},
				},
				Value: 45,
				Right: &BinaryTree{
					Left:  nil,
					Value: 60,
					Right: nil,
				},
			},
		},
		Value: 40,
		Right: &BinaryTree{
			Left:  nil,
			Value: 50,
			Right: nil,
		},
	}
)
