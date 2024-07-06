package main 
import (
	"fmt"
)

func main() {
	var inputString string 
	var reverseString string 
	fmt.Println("Program to check palindrome: ")
	fmt.Println("Input a String: ")
	fmt.Scan(&inputString)
	for i:=len(inputString)-1;i>=0;i-- {
		reverseString += string(inputString[i])
	}
	fmt.Printf("%s\n", reverseString)
	if inputString == reverseString {
		fmt.Printf("Input string %v is a palindrome\n", inputString)
	} else {
		fmt.Printf("Input string %v is not a palindrome\n", inputString)
	}
}