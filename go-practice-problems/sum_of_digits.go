package main 
import "fmt"

func main() {
	var inputNumber int
	var quotient int = 1
	var remainder int
	var digitSum int = 0
	fmt.Println("Program to print sum of digits: ")
	fmt.Println("Enter a number: ")
	fmt.Scan(&inputNumber)
	for quotient > 0 {
		quotient = inputNumber/10 
		remainder = inputNumber % 10 
		digitSum = digitSum + remainder
		inputNumber = quotient
	}
	fmt.Printf("Sum of digits entered is: %v\n", digitSum)
}
