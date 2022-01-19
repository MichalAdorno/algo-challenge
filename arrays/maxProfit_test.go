package arrays

import "testing"

func Test_CalculateHighestProfit(t *testing.T) {
	var tests = []struct {
		prices    []int
		buyPrice  int
		sellPrice int
		maxProfit int
	}{
		{
			prices:    []int{310, 315, 275, 295, 260, 270, 290, 230, 255, 250, 285},
			buyPrice:  230,
			sellPrice: 285,
			maxProfit: 55,
		},
		{
			prices:    []int{310, 315, 275, 295, 260, 270, 290, 230, 255},
			buyPrice:  260,
			sellPrice: 290,
			maxProfit: 30,
		},
		{
			prices:    []int{},
			buyPrice:  0,
			sellPrice: 0,
			maxProfit: 0,
		},
		{
			prices:    []int{310},
			buyPrice:  0,
			sellPrice: 0,
			maxProfit: 0,
		},
		{
			prices:    []int{310, 315},
			buyPrice:  310,
			sellPrice: 315,
			maxProfit: 5,
		},
		{
			prices:    []int{310, 315, 275},
			buyPrice:  310,
			sellPrice: 315,
			maxProfit: 5,
		},
		{
			prices:    []int{310, 315, 275, 295},
			buyPrice:  275,
			sellPrice: 295,
			maxProfit: 20,
		},
		{
			prices:    []int{310, 315, 275, 295, 260},
			buyPrice:  275,
			sellPrice: 295,
			maxProfit: 20,
		},
		{
			prices:    []int{310, 315, 275, 295, 260, 210},
			buyPrice:  275,
			sellPrice: 295,
			maxProfit: 20,
		},
		{
			prices:    []int{310, 315, 275, 295, 260, 270},
			buyPrice:  275,
			sellPrice: 295,
			maxProfit: 20,
		},
		{
			prices:    []int{310, 315, 275, 295, 260, 270, 290},
			buyPrice:  260,
			sellPrice: 290,
			maxProfit: 30,
		},
		{
			prices:    []int{310, 315, 275, 295, 260, 270, 210},
			buyPrice:  275,
			sellPrice: 295,
			maxProfit: 20,
		},
		{
			prices:    []int{310, 315, 275, 295, 260, 270, 210, 250},
			buyPrice:  210,
			sellPrice: 250,
			maxProfit: 40,
		},
		{
			prices:    []int{310, 315, 275, 295, 260, 270, 210, 250, 190, 230},
			buyPrice:  210,
			sellPrice: 250,
			maxProfit: 40,
		},
		{
			prices:    []int{310, 315, 275, 295, 260, 270, 210, 250, 190, 230, 180, 220},
			buyPrice:  210,
			sellPrice: 250,
			maxProfit: 40,
		},
		{
			prices:    []int{310, 315, 275, 295, 260, 270, 210, 250, 190, 230, 180, 220, 170, 180, 160, 210},
			buyPrice:  160,
			sellPrice: 210,
			maxProfit: 50,
		},
		{
			prices:    []int{400, 310, 315, 275, 295, 260, 270, 210, 250, 190, 230, 180, 220, 170, 180, 160, 210},
			buyPrice:  160,
			sellPrice: 210,
			maxProfit: 50,
		},
		{
			prices:    []int{400, 310, 315, 275, 295, 260, 270, 210, 250, 190, 230, 180, 220, 170, 180, 160, 210, 900},
			buyPrice:  160,
			sellPrice: 900,
			maxProfit: 740,
		},
		{
			prices:    []int{400, 410, 420, 430, 440, 450, 900},
			buyPrice:  400,
			sellPrice: 900,
			maxProfit: 500,
		},
		{
			prices:    []int{500, 410, 320, 230, 140, 50, 600},
			buyPrice:  50,
			sellPrice: 600,
			maxProfit: 550,
		},
		{
			prices:    []int{500, 410, 320, 230, 140, 50},
			buyPrice:  0,
			sellPrice: 0,
			maxProfit: 0,
		},
		{
			prices:    []int{500, 410, 320, 230, 140, 50},
			buyPrice:  0,
			sellPrice: 0,
			maxProfit: 0,
		},
	}
	for _, test := range tests {
		if buyPrice, sellPrice, highestProfit := CalculateMaxProfit(test.prices); buyPrice != test.buyPrice || sellPrice != test.sellPrice || highestProfit != test.maxProfit {
			t.Errorf("For price time series: [%v]\n Expected maxProfit is: [%d]\n But got: [%d]\n",
				test.prices, test.maxProfit, highestProfit)
			t.Errorf("buyPrice=[%v/t:%v], sellPrice=[%v/t:%v]\n",
				buyPrice, test.buyPrice, sellPrice, test.sellPrice)
		}
	}
}
