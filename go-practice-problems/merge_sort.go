// implement merge sort

package main 
import "fmt"

func main() {
	var inputArray = []int{10, 5, 3, 4}
	var C =[]int{}
	// var A =[]int{2, 4}
	// var B =[]int{5, 6}	
	// divideArray(inputArray)
	C = mergeSort(inputArray)
	// C = mergeArrays(A,B)
	fmt.Printf("Merged Array is: %v\n", C)
}

func mergeSort(inputArray []int) []int{
	var mid int 
	var A []int 
	var B []int
	if len(inputArray)==1 {
		return inputArray
	}
	mid = len(inputArray)/2

	A = mergeSort(inputArray[0:mid])
	// fmt.Printf("A: %v\n", A)
	B = mergeSort(inputArray[mid:len(inputArray)])
	// fmt.Printf("B: %v\n", B)
	return mergeArrays(A, B)
}



func mergeArrays(A []int, B []int) []int{
	var i int = 0
	var j int = 0
	var sortedArray []int 
	// var k int 
	for ;i<len(A) && j<len(B); {
		if A[i] < B[j] {
			sortedArray = append(sortedArray, A[i])
			i = i + 1 
			// k = k+1
		} else if B[j] < A[i] {
			sortedArray = append(sortedArray, B[j])
			// k = k + 1 
			j = j + 1
		}
	}
	if i == len(A) {
		for j < len(B) {
			sortedArray = append(sortedArray, B[j])
			// k=k+1 
			j=j+1
		}
	} else {
		for i < len(A) {
			sortedArray = append(sortedArray, A[i])
			// k=k+1
			i=i+1
		}
	}
	// fmt.Printf("Sorted Array: %v", sortedArray)
	return sortedArray
}