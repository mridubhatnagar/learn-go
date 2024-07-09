package main
import "fmt"

func main() {
	var a = [3][2]int{{1,2}, {2,3}, {4,5}}
	var b = [3][2]int{{5,6}, {7,8}, {9,10}}
	var rows int = 3
	var cols int = 2
	var c[3][2]int 
	for i:=0;i<rows;i++ {
		for j:=0;j<cols;j++ {
			c[i][j] = a[i][j] + b[i][j]
		}
	}
	for i:=0; i<rows;i++ {
		for j:=0;j<cols;j++ {
			fmt.Printf("%d ", c[i][j])
		}
		fmt.Print("\n")
	}  
}