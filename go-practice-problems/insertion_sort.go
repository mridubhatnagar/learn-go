package main 

import "fmt"

func main() {
	var inputArray = [7]int{9, 12, 14, 15, 6, 13, 1}
	var result [7]int 
	result = insertionSort(inputArray)
	fmt.Printf("Sorted array is: %v", result)
}

func insertionSort(inputArray [7]int) [7]int {
	var j int
	var temp int
	for i:=0;i<=len(inputArray)-2;i++ {
		j = i+1
		for ;j>0&&inputArray[j] < inputArray[j-1]; j-- {
			temp = inputArray[j]
			inputArray[j] = inputArray[j-1]
			inputArray[j-1] = temp 
		}
	}
	return inputArray
}