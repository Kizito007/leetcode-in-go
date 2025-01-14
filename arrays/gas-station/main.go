package main

import "fmt"

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}

	fmt.Println(canCompleteCircuit(gas, cost))
}

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	remainingFuel, globalFuelRemaining, start := 0, 0, 0
	for i := 0; i < n; i++ {
		globalFuelRemaining += gas[i] - cost[i]
		remainingFuel += gas[i] - cost[i]
		if remainingFuel < 0 {
			start = i + 1
			remainingFuel = 0
		}
	}
	if globalFuelRemaining < 0 {
		return -1
	}
	return start
}
