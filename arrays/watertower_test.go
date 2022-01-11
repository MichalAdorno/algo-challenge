package arrays

import (
	"testing"
)

func Test_lowToHighAndEqual(t *testing.T) {
	var tests = []struct {
		watertower  []int
		wateramount int
	}{
		{
			watertower:  []int{3, 1},
			wateramount: 0,
		},
		{
			watertower:  []int{3, 1},
			wateramount: 0,
		},
		{
			watertower:  []int{1},
			wateramount: 0,
		},
		{
			watertower:  []int{3, 1, 2, 1, 4},
			wateramount: 5,
		},
		{
			watertower:  []int{3, 1, 2, 1, 4, 1, 2, 1, 4, 5},
			wateramount: 13,
		},
		{
			watertower:  []int{1, 2, 1, 2, 1, 2, 1, 2},
			wateramount: 3,
		},
		{
			watertower:  []int{1, 2, 1},
			wateramount: 0,
		},
		{
			watertower:  []int{2, 1, 2},
			wateramount: 1,
		},
		{
			watertower:  []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			wateramount: 0,
		},
		{
			watertower:  []int{9, 8, 7, 6, 5, 4, 3, 2, 7},
			wateramount: 0,
		},
		{
			watertower:  []int{9, 8, 7, 6, 5, 4, 3, 2, 3},
			wateramount: 0,
		},
		{
			watertower:  []int{9, 8, 7, 6, 5, 4, 3, 2, 9},
			wateramount: 28,
		},
	}
	for _, test := range tests {
		if res := lowToHighAndEqual(test.watertower); res != test.wateramount {
			t.Errorf("For \"lowTohigh\" watertower %v expected water amount is [%d] but got: [%d]", test.watertower, test.wateramount, res)
		}
	}
}

func Test_highToLow(t *testing.T) {
	var tests = []struct {
		watertower  []int
		wateramount int
	}{
		{
			watertower:  []int{3, 1},
			wateramount: 0,
		},
		{
			watertower:  []int{3, 1},
			wateramount: 0,
		},
		{
			watertower:  []int{1},
			wateramount: 0,
		},
		{
			watertower:  []int{3, 1, 2, 1, 4},
			wateramount: 0,
		},
		{
			watertower:  []int{3, 1, 2, 1, 4, 1, 2, 1, 4, 5},
			wateramount: 0,
		},
		{
			watertower:  []int{1, 2, 1, 2, 1, 2, 1, 2},
			wateramount: 0,
		},
		{
			watertower:  []int{1, 2, 1},
			wateramount: 0,
		},
		{
			watertower:  []int{2, 1, 2},
			wateramount: 0,
		},
		{
			watertower:  []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			wateramount: 0,
		},
		{
			watertower:  []int{9, 8, 7, 6, 5, 4, 3, 2, 3},
			wateramount: 1,
		},
		{
			watertower:  []int{9, 8, 7, 6, 5, 4, 3, 2, 5},
			wateramount: 6,
		},
		{
			watertower:  []int{9, 8, 7, 6, 5, 4, 3, 2, 9},
			wateramount: 0,
		},
		{
			watertower:  []int{9, 8, 7, 6, 5, 4, 3, 2, 4, 3, 4},
			wateramount: 4,
		},
		{
			watertower:  []int{6, 1, 9, 8, 7, 6, 5, 4, 3, 2, 4, 3, 4},
			wateramount: 4,
		},
	}
	for _, test := range tests {
		if res := highToLow(test.watertower); res != test.wateramount {
			t.Errorf("For \"highToLow\" watertower %v expected water amount is [%d] but got: [%d]", test.watertower, test.wateramount, res)
		}
	}
}

func Test_WaterAmountCalculation(t *testing.T) {
	var tests = []struct {
		watertower  []int
		wateramount int
	}{
		{
			watertower:  []int{3, 1},
			wateramount: 0,
		},
		{
			watertower:  []int{3, 1},
			wateramount: 0,
		},
		{
			watertower:  []int{1},
			wateramount: 0,
		},
		{
			watertower:  []int{3, 1, 2, 1, 4},
			wateramount: 5,
		},
		{
			watertower:  []int{3, 1, 2, 1, 4, 1, 2, 1, 4, 5},
			wateramount: 13,
		},
		{
			watertower:  []int{1, 2, 1, 2, 1, 2, 1, 2},
			wateramount: 3,
		},
		{
			watertower:  []int{1, 2, 1},
			wateramount: 0,
		},
		{
			watertower:  []int{2, 1, 2},
			wateramount: 1,
		},
		{
			watertower:  []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			wateramount: 0,
		},
		{
			watertower:  []int{9, 8, 7, 6, 5, 4, 3, 2, 3},
			wateramount: 1,
		},
		///////////////////////
		{
			watertower:  []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			wateramount: 0,
		},
		{
			watertower:  []int{9, 8, 7, 6, 5, 4, 3, 2, 3},
			wateramount: 1,
		},
		{
			watertower:  []int{9, 8, 7, 6, 5, 4, 3, 2, 5},
			wateramount: 6,
		},
		{
			watertower:  []int{9, 8, 7, 6, 5, 4, 3, 2, 4, 3, 4},
			wateramount: 4,
		},
		///////////////////////
		{
			watertower:  []int{6, 1, 9, 8, 7, 6, 5, 4, 3, 2, 4, 3, 4},
			wateramount: 9,
		},
		{
			watertower:  []int{9, 8, 7, 6, 5, 4, 3, 2, 9},
			wateramount: 28,
		},
		{
			watertower:  []int{3, 2, 1, 2, 3, 4, 5, 6, 7, 8, 7, 6, 5, 4, 3, 1, 5, 4, 3, 2, 1, 2},
			wateramount: 12,
		},
	}
	for _, test := range tests {
		if res := MaxWater(test.watertower); res != test.wateramount {
			t.Errorf("For \"MaxWater\" watertower %v expected water amount is [%d] but got: [%d]", test.watertower, test.wateramount, res)
		}
	}
}
