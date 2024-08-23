package main 
import (
	"fmt"
)

func main() {
	var inputString string = "BOBBY RAN AFTER MADAM MONICA SABASTIAN"
	var fixedLength int = len(inputString)
	var minLength int = 2 
	var maxLength int = len(inputString)
	
	
	// checks palindrome for fixed length N.
	for i:=0;i+fixedLength<=len(inputString);i++{
		result := checkPalindromeIterative(inputString[i:i+fixedLength], fixedLength)
		
		if result {
			fmt.Printf("%s is Palindrome: %v\n", inputString[i:i+fixedLength], result)	
		}
	}
	// checks palindrome upto length N
	for minLength <= maxLength {
		var result bool 
		for i:=0;i+minLength<=len(inputString);i++{
			result = checkPalindromeIterative(inputString[i:i+minLength], minLength)
			if result {
				fmt.Printf("%s Palindrome: %v\n", inputString[i:i+minLength], result)
			} 
		}
		minLength += 1
	}
}


func checkPalindrome(s string, length int) bool{
	var mid int 
	mid = length/2
	var forward string = ""
	var backward string = ""
	for i:=0;i<mid;i++ {
		forward += string(s[i])
	}
	for j:=length-1;j>mid;j-- {
		backward += string(s[j])
	}
	if forward == backward {
		return true 
	} else {
		return false 
	}
}

// two pointer approach
func checkPalindromeIterative(s string, length int) bool {
	var i int = 0
	var j int
	for j=length-1; j>=i ; j-- {
		if string(s[i]) != string(s[j]) {
			return false
		}
		i += 1
	}
	return true 
}