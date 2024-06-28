// Q. Display the fibonacci sequence up to the total number of terms that has been input by the user.
//    Fibonacci sequence is : 0 1  1  2  3  5  8 13  21 ....
//    So if we enter the total number of terms as 5, then the first 5 fibonacci numbers are to be printed

package main 
import "fmt"

func main() {
	fmt.Println("Program to Generate Fibonnaci Sequence: ")
	var numCount int
	fmt.Println("Enter number count: ")
	fmt.Scan(&numCount)
	firstNum := 0 
	secondNum := 1
	var thirdNum int  

	if numCount >= 1 {
		fmt.Printf("%v ", firstNum)
	} 
	if numCount >= 2 {
		fmt.Printf("%v ", secondNum)
	}  
	if numCount >= 3 {
		for i:=3; i<=numCount; i++ {		
			thirdNum = firstNum + secondNum
			fmt.Printf("%v ", thirdNum)
			firstNum = secondNum
			secondNum = thirdNum  
			}
	} 	
	}
	
	
