// print numbers from 1 to N using recursion
package main

import "fmt"

func main() {
	var inputNumber int = 10
	var currentNumber int = 1
	_ = printNtimes(inputNumber, currentNumber)
}

func printNtimes(N int, i int) int{
	if N == 0 {
		return 1
	} else {
		fmt.Printf("%d ", i)
		return printNtimes((N - 1), i+1)
	}
}