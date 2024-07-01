// Q. Write a program that would reverse the values in the array. So element at zero should go to the last location and vice versa for each corresponding pairs.
package main
import "fmt"

func main() {
	var numbers[6]int 
	var number int 
	var temp int 
	var j int = len(numbers)-1
	for i:=0;i<len(numbers);i++ {
		fmt.Scan(&number)
		numbers[i] = number
	}
	for i:=0; i < len(numbers)/2; i++ {
		temp = numbers[i]
		numbers[i] = numbers[j]
		numbers[j] = temp 
		j = j-1
	}
	fmt.Printf("%v\n", numbers)
	} 
