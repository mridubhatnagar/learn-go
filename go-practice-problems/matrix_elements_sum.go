package main

import "fmt"

func main() {
	var array2D = [3][3]int{{1,2,3},{4,5,6},{7,8,9}}
	var rows int = 3 
	var cols int = 3
	var sum int = 0 
	for i:=0;i<rows;i++ {
		for j:=0;j<cols;j++ {
			sum += array2D[i][j]
		}
	}
	fmt.Printf("Sum of elements: %d\n", sum)
}