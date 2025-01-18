package main

import "fmt"

func main() {
	fmt.Println(romanToInt("XXVII"))
	fmt.Println(romanToInt("IV"))      // Output: 4
	fmt.Println(romanToInt("MCMXCIV")) // Output: 1994
}

func romanToInt(s string) int {
	numerals := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	total := 0
	n := len(s)

	for i := 0; i < n; i++ {
		// If the current numeral is less than the next numeral, subtract it
		if i < n-1 && numerals[s[i]] < numerals[s[i+1]] {
			fmt.Println(numerals[s[i]])
			total -= numerals[s[i]]
		} else {
			total += numerals[s[i]]
		}
	}

	return total
}
