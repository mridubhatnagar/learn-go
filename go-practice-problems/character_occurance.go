package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var inputString string
	var ctr int 
	// var char string 
	// var possibleChars string = "abcdefghijklmnopqrstuvwxyz"
	reader := bufio.NewReader(os.Stdin)
	inputString, _ = reader.ReadString('\n')
	for i:='a'; i <= 'z'; i++ {
		ctr = 0
		// char = string(rune(inputString[i]))
		for j:=0; j < len(inputString); j++ {
			//if string(rune(inputString[j])) == char {
			if rune(inputString[j]) == i {
				ctr += 1
			}	
		}
		if ctr >= 1 {
			fmt.Printf("character: %c, occurance: %v\n", i, ctr)
		}
	}
}