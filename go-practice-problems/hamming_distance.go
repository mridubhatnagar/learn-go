package main 
import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	var firstString string 
	var secondString string 
	var ctr int 
	fmt.Println("Program to calculate hamming distance between two strings: ")
	fmt.Println("Input first string: ")
	reader := bufio.NewReader(os.Stdin)
	firstString, _ = reader.ReadString('\n')
	fmt.Println("Input second string: ")
	secondString, _ = reader.ReadString('\n')
	for i:=0;i<len(firstString);i++ {
		if firstString[i] != secondString[i] {
			ctr += 1
		}
	}
	fmt.Printf("Hamming Distance: %d\n", ctr)
}