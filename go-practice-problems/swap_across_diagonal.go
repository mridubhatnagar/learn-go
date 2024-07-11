package main
import "fmt"

func main() {
	var a = [4][4]int{{1,2,3,4},{5,6,7,8}, {9,10,11,12}, {13,14,15,16}}
	var temp int 
	var rows int = 4
	var cols int = 4
	for i:=0;i<rows;i++ {
		for j:=0;j<cols;j++ {
			if i<j {
				temp = a[i][j]
				a[i][j] = a[j][i]
				a[j][i] = temp 
			}
		}
	}
	for i:=0;i<rows;i++ {
		for j:=0;j<cols;j++ {
			fmt.Printf("%d ", a[i][j])
		}
		fmt.Printf("\n")
	}
}