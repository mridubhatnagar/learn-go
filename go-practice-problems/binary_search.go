package main 
import "fmt"

func main() {
	var a = [5]int{23, 35, 44, 57, 96}
	var low int = 0
	var high int = len(a) - 1
	var mid int 
	var target int 
	fmt.Println("Enter element to search: ")
	fmt.Scan(&target)
	for low <= high {
		mid = (low+high)/2
		if target == a[mid] {
			fmt.Printf("Element found at %v\n", mid)
			break
		} else if target < a[mid] {
			high = mid-1
		} else if target > a[mid] {
			low = mid+1
		}
		if low > high {
			fmt.Printf("Element not found: -1\n")
			break
		}
	}
}