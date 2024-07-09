// Q. Sort an array of strings
package main

import "fmt"

func main() {
	var langs = [5]string{"python", "haskell", "rust", "go", "c"}
	var temp string 
	var smallest_index int 
	for i:=0;i < len(langs)-1;i++ {
		smallest_index = i
		for j:=i+1;j < len(langs);j++ {
			if langs[j] < langs[smallest_index] {
				smallest_index = j 
			} 
		}
		temp = langs[smallest_index]
		langs[smallest_index] = langs[i]
		langs[i] = temp 
	}
	fmt.Printf("Alphabetically sorted array is: %v\n", langs)
}
