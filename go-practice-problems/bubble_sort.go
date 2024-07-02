// logic and implementation bubble sort.
// logic - Find maximum and place it at the end of array.
// from remaining array find max and place at second last and so on. 

package main
import "fmt"

func main() {
	var a = [5]int{32, 44, 23, 12, 6}
	var temp int 
	for i:=len(a)-1;i>=0;i-- {
		for j:=0;j<=len(a)-2;j++ {
			if a[j] > a[j+1] {
				temp = a[j+1]
				a[j+1] = a[j]
				a[j] = temp 
			}
		}
	}
	fmt.Printf("Sorted array using bubble sort: %v", a)
}