// logic for selection sort 

package main 
import "fmt"

func main() {
	var numbers=[6]int{23, 32, 4, 5, 55, 21}
	var temp int 
	var smallest_index int 
	for i:=0;i<=len(numbers)-2;i++ {
		smallest_index = i 
		for j:=i+1;j<=len(numbers)-1;j++ {
			if numbers[j] < numbers[smallest_index] {
				smallest_index = j 
			}
		}
		temp = numbers[smallest_index]
		numbers[smallest_index] = numbers[i]
		numbers[i] = temp
	}
	fmt.Printf("Sorted array: %v\n", numbers)
}