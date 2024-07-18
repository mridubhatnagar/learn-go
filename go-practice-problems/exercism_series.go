// solution to exercism problem series

package main
import "fmt" 

func main() {
	var allSubStrings = []string{}
	allSubStrings = All(3, "63915")
	fmt.Printf("All possible substrings are %v\n", allSubStrings)
	var firstSubString string 
	firstSubString = UnsafeFirst(3, "63915")
	fmt.Printf("First substring is %v\n", firstSubString)


}

func All(n int, s string) []string {
    var subStrings = []string{}
    var subString string
	for i:=0;i<=len(s)-n;i++ {
        subString = s[i:i+n]
        subStrings = append(subStrings, subString)
    }
    return subStrings
}
func UnsafeFirst(n int, s string) string {
    var subString string
	for i:=0;i<=len(s)-n;i++ {
        subString = s[i:i+n]
        break
    }
    return subString
}