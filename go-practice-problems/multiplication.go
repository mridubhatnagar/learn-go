package main 

import "fmt"

func main() {
	var inputNumber int
	fmt.Println("Generate multiplication table for: ")
	fmt.Scan(&inputNumber)
	for i:=1; i<=10; i++ {
		result := inputNumber * i 
		fmt.Printf("%v x %v = %v\n", inputNumber, i, result)
	}

}