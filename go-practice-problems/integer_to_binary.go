package main 
import "fmt"

func main() {
	var inputNumber int 
	fmt.Println("Program to convert Integer to binary")
	fmt.Println("Input the number: ")
	fmt.Scan(&inputNumber)
	var output string = ""
	var remainder int 
	var quotient int
	for inputNumber > 0 {
		remainder = inputNumber%2
		if remainder == 1 {
			output = "1" + output
		} else {
			output = "0" + output
		}
		quotient = inputNumber/2
		inputNumber = quotient
	}
	fmt.Printf("%s", output)
}