package main 
import "fmt"

func main() {
	var inputNumber int 
	fmt.Println("Program to convert Integer to binary")
	fmt.Println("Input the number: ")
	fmt.Scan(&inputNumber)
	var output = []int{}
	var remainder int 
	var quotient int 
	for inputNumber > 0 {
		remainder = inputNumber%2 
		quotient = inputNumber/2
		inputNumber = quotient
		output = append(output, remainder)
	}
	fmt.Printf("%v", output)
}