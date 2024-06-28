package main

import "fmt"

func main() {
	var inputNumber int
	fmt.Println("Program sum of N numbers: ")
	fmt.Println("Enter the value for N: ")
	fmt.Scan(&inputNumber)
	var totalSum int
	totalSum = 0 
	for i:=1; i<=inputNumber; i ++ {
		totalSum = totalSum + i 
	}
	fmt.Printf("Sum of first %v numbers is %v\n", inputNumber, totalSum)
}