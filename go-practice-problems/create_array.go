// Q. Write a program that has an array of 5 integers containing some data. The program should display all the values that occur only once in the array.
package main 
import "fmt"

func main() {
	var numbers[5]int
	var number int 
	var duplicate bool
	var j int 
	var arrayLen int = len(numbers)
	for i:=0; i< len(numbers); i++ {
		fmt.Scan(&number)
		numbers[i] = number
	}
	fmt.Printf("Array elements are %v\n", numbers)
	fmt.Println("Unique elements are: ")
	for i:=0;i<len(numbers);i++ {
		duplicate = false 
		for j=0; j<len(numbers); j++ {
			if numbers[i] == numbers[j] && i != j {
				duplicate = true  
				break
			} 
		}
		if duplicate == false && j == arrayLen {
			fmt.Printf("%v\n", numbers[i])
		}
	}
}