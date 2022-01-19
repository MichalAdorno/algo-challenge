package arrays

func CalculateMaxProfit(prices []int) (int, int, int) {
	if len(prices) < 2 {
		return 0, 0, 0
	}
	prevBuy := prices[0]
	prevSell := prices[0]
	prevProfit := 0

	buy := prevBuy
	sell := prevSell
	profit := 0

	for i := 1; i < len(prices); i++ {
		newProfit := prices[i] - buy
		if newProfit >= profit {
			if prevProfit < profit {
				prevBuy = buy
				prevSell = prices[i]
				prevProfit = newProfit
			}
			sell = prices[i]
			profit = newProfit

		} else if prices[i] < buy {
			if prevProfit < profit {
				prevBuy = buy
				prevSell = sell
				prevProfit = profit
			}

			buy = prices[i]
			sell = prices[i]
			profit = newProfit

		}

	}
	if profit <= prevProfit {
		buy = prevBuy
		sell = prevSell
		profit = prevProfit

	}
	if profit <= 0 {
		buy = 0
		sell = 0
		profit = 0

	}

	return buy, sell, profit
}
