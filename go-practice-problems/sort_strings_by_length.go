// Q. In an array of strings ( pre defined and initialised ), sort the array on basis of length of the strings. So shortest string first.
package main 
import "fmt"

func main() {
	// can 2 strings be of same length? 
	var langs = [6]string{"python", "go", "c", "rust", "haskell"}
	var smallest_index int 
	var temp string
	for i:=0;i<=len(langs)-2;i++ {
		smallest_index = i 
		for j:=i+1; j <=len(langs)-1;j++ {
			if len(langs[j]) < len(langs[smallest_index]) {
				smallest_index = j 
			}
		}
		temp = langs[smallest_index]
		langs[smallest_index] = langs[i]
		langs[i] = temp 
	}
	fmt.Printf("Sorted array: %v\n", langs)
}
