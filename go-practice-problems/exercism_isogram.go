// solution to exercism isogram exercise problem

package main 
import "fmt"

func main() {
	var inputString string = "downstream"
	var count int 
	var isIsogram bool=true 
	// var isIsogram bool = false 
	for i:='a';i<='z'; i++ {
		count = 0
		for _, value := range inputString {
			if i == value {
				count += 1
			}
			if count > 1 {
				isIsogram = false 
			}
		}
	}
	if isIsogram == false {
		fmt.Printf("inputString %s is not Isogram\n", inputString)
	} else {
		fmt.Printf("inputString %s is a Isogram\n", inputString)
	}
}