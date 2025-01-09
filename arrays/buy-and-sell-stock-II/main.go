package main

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}

	println(buyAndSellStock(prices))
}

func buyAndSellStock(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}
	return maxProfit
}
