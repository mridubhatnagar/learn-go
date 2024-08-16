package main 

import "fmt"

func main() {
	var inputNumber int = 5
	result := factorialIterative(inputNumber)
	fmt.Printf("Factorial of %d is %d\n", inputNumber, result)
	resultr := factorialRecursive(inputNumber)
	fmt.Printf("Factorial of %d is %d\n", inputNumber, resultr)
}

func factorialIterative(n int) int {
	var result int = 1
	for i:=0;i<=n-1;i++ {
		result = result * (n-i)
	}
	return result
}

func factorialRecursive(n int) int {
	if n == 1{
		return 1
	} else {
		return n * factorialRecursive(n-1)
	}
}