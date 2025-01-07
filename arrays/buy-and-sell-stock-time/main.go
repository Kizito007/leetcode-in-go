package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}

	fmt.Println(buyAndSellStock(prices))
}

func buyAndSellStock(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}
	return maxProfit
}
