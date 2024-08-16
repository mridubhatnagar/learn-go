package main

import "fmt"

func main() {
	var inputNumber int = 5 
	var result int 
	var resultr int 
	result = sumOfNIterative(inputNumber)
	fmt.Printf("Sum of %d integers (iterative) is %d\n", inputNumber, result)
	resultr = sumOfNRecursive(inputNumber)
	fmt.Printf("Sum of %d integers (recursive) is %d\n", inputNumber, resultr)

}

func sumOfNIterative(n int) int {
	var sum int = 0
	for i:=0;i<=n;i++ {
		sum = sum + i 
	}
	return sum 
}

func sumOfNRecursive(n int) int {
	if n == 0 {
		return 0
	} else {
		return n + sumOfNRecursive(n-1)
	}
}