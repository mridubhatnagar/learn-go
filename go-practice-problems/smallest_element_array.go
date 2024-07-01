package main 
import "fmt"

func main() {
	var numbers[5]int 
	var smallest int 
	var number int 
	fmt.Println("Enter elements in a array")
	for i:=0;i<len(numbers);i++ {
		fmt.Scan(&number)
		numbers[i] = number
	}
	smallest = numbers[0]
	for i:=0;i<len(numbers);i++ {
		if numbers[i] < smallest {
			smallest = numbers[i]
		}
	}
	fmt.Printf("Smallest element is %v\n", smallest)
}