package main

import "fmt" 

func main() {
	var inputNumber uint
	fmt.Println("Check if a number is Even or Odd:")
	fmt.Printf("Enter a number (%T): \n", inputNumber)
	fmt.Scan(&inputNumber)
	if inputNumber%2==0 {
		fmt.Printf("%d is even\n", inputNumber)
	} else {
		fmt.Printf("%d is odd\n", inputNumber)
	}
}
