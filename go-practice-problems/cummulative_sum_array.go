// Q. Write a program to display cumulative sum of values.
//    So if the numbers in the array are 5,7,1,4,8
//    Then the output should be:
//    5  5
//    7  12
//    1  13
//    4  17
//    8  25
//    Notice that the second column is the cumulative sum of numbers as they occur in array.
package main
import "fmt"

func main() {
	var sum int 
	var numbers[5]int 
	var number int 
	fmt.Printf("Enter elements in array: ")
	for i:=0; i<len(numbers);i++ {
		fmt.Scan(&number)
		numbers[i] = number
	}
	for i:=0;i<len(numbers);i++ {
		sum = 0
		for j:=0; j<=i;j++ {
			sum = sum + numbers[j]
		}
		fmt.Printf("%v %v\n", numbers[i], sum)
	}
}