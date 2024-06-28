// Q. Input a count of numbers, and then input those many numbers. At the end display the largest, smallest, sum and average of the numbers entered.

package main 
import "fmt"

func main() {
	var count uint 
	fmt.Println("Enter count of numbers: ")
	fmt.Scan(&count)
	var largest uint = 0
	var smallest uint = 1
	var sum uint = 0
	var average float64
	var userInput uint 
	for i:=uint(1); i <= count; i++ {
		fmt.Printf("Input Number %v:" , i)
		fmt.Scan(&userInput)
		if userInput > largest {
			largest = userInput
		} 
		if userInput < smallest {
			smallest = userInput
		}
		sum = sum + userInput
		 
	}
	average = float64(sum)/float64(count)

	fmt.Printf("Count of numbers entered: %v\n", count)
	fmt.Printf("Largest number is: %v\n", largest)
	fmt.Printf("Smallest number is: %v\n", smallest)
	fmt.Printf("Sum of all numbers is: %v\n", sum)
	fmt.Printf("Average of all numbers is: %v\n", average)
}