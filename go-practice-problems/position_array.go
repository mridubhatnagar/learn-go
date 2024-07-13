package main 
import (
	"fmt"
	"strings"
)

func main() {
	var inputString string = "ABCDEFGHIJKLDEF"
	var subString string = "F"
	var positions[]int
	positions = getPositions(inputString, subString)
	fmt.Printf("%v\n", positions)
	
}

func getPositions(inputString string, subString string) []int {
	var remainingString string 
	var index int 
	var prevIndex int 
	var ctr int = 1 
	var positions []int 

	for {
		index = strings.Index(inputString, subString)
		remainingString = inputString[index+1:]
		inputString = remainingString
		if index == -1 {
			break
		}
		if ctr > 1 {
			index = prevIndex + index 
		}
		positions = append(positions, index)
		prevIndex = index 
		ctr += 1
	}
	return positions
}