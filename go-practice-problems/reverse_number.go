// Q. Input a number and reverse it and display the reversed number
//    So if number entered was 458
//    The reversed number displayed should be 854
//    We can enter a number of any valid length

package main 

import "fmt"

func main() {
	fmt.Println("Program to Reverse a Number")
	fmt.Println("Enter a number: ")
	var inputNumber int
	fmt.Scan(&inputNumber)
	var quotient int = 1
	var remainder int
	var new_number int 
	for quotient > 0 {
		quotient = inputNumber/10
		remainder = inputNumber%10
		new_number = new_number * 10 + remainder
		inputNumber = quotient
	}
	fmt.Printf("Reversed number is: %v \n", new_number)
}