package main 
import "fmt"

func main() {
	var inputNumber int 
	var quotient int = 1
	var remainder int
	var sumOfCubes int = 0 
	fmt.Println("Program to check if input number is Amstrong: ")
	fmt.Println("Enter a Number: ")
	fmt.Scan(&inputNumber)
	var originalNumber int 
	originalNumber = inputNumber
	for quotient > 0 {
		quotient = inputNumber/10
		remainder = inputNumber % 10 
		sumOfCubes = sumOfCubes + remainder*remainder*remainder
		inputNumber = quotient
	}
	if originalNumber == sumOfCubes {
		fmt.Printf("%v is a Amstrong Number\n", originalNumber)
	} else {
		fmt.Printf("%v is not a Amstrong Number\n", originalNumber)
	}
}