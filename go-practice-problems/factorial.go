package main 
import "fmt"

func main() {
	var inputNumber int
	fmt.Scan(&inputNumber) 
	result := factorial(inputNumber)
	fmt.Printf("Factorial of %d is %v\n", inputNumber, result)

}

func factorial(inputNumber int) int {
	var r int
	var result int = 1
	for r < inputNumber {
		result *= inputNumber - r
		r += 1
	}
	return result  
}