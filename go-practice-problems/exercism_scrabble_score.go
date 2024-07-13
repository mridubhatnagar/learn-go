// solution to exercism scrabble score go problem

package main

import (
	"fmt"
	"strings"
)

func main() {
	var inputString string
	fmt.Scan(&inputString)
	inputString = strings.ToUpper(inputString)
	var score int = 0
	for _, value := range inputString {
		if value == 'A' || value=='E' || value=='I' || value=='O' || value=='U' ||
		value == 'L' || value=='N' || value=='R' || value=='S' || value=='T' {
			score += 1
		} else if value == 'D' || value == 'G' {
			score += 2
		} else if value == 'B' || value == 'C' || value == 'M' || value == 'P' {
			score += 3
		} else if value == 'F' || value == 'H' || value == 'V' || value == 'W' || value == 'Y' {
			score += 4
		} else if value == 'K' {
			score += 5
		} else if value == 'J' || value == 'X' {
			score += 8
		} else if value == 'Q' || value == 'Z' {
			score += 10
		}
	}
	fmt.Printf("Score: %d\n", score)
}