package binaryTrees

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_BuildMinimalTree(t *testing.T) {
	var tests = []struct {
		values         []int
		expectedValues []int
	}{
		{
			values:         []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 90},
			expectedValues: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 90},
		},
	}
	for _, test := range tests {
		res := InOrder(BuildMinimalTree(test.values))
		fmt.Println("***")
		fmt.Println(res)
		fmt.Println(test.expectedValues)
		fmt.Println("***")
		bt := BuildMinimalTree(test.values)
		fmt.Println(bt)
		fmt.Println(bt.Left)
		fmt.Println(bt.Right)

		if res == nil || reflect.DeepEqual(test.expectedValues, res) {
			t.Errorf("Got [%v] but the expected order is [%v]",
				test.values, test.expectedValues)
		}
	}
}
