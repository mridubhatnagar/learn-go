// Q. Write a program that would input a number and search for it in the array. The program should display the location where it was found, otherwise it should display -1
package main 
import "fmt"

func main() {
	var number int
	var numbers[5] int 
	var searchNum int 
	var found int = -1
	for i:=0;i<len(numbers);i++ {
		fmt.Scan(&number)
		numbers[i]=number
	}
	fmt.Println("Enter element to search: ")
	fmt.Scan(&searchNum)
	for index, value := range numbers {
		if searchNum == value {
			found = 0
			fmt.Printf("Found at %v value %v: \n", index, value)
			break 
		} 
	}
	if found == -1 {
		fmt.Printf("%v\n", found)
	} 
}