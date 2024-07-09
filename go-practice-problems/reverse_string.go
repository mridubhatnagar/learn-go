package main 
import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var inputString string 
	var reverseString string 
	reader := bufio.NewReader(os.Stdin)
	inputString, _ = reader.ReadString('\n')
	for i:=len(inputString)-1; i>=0;i-- {
		reverseString += string(inputString[i])
	}
	fmt.Printf("%s\n", reverseString)
}