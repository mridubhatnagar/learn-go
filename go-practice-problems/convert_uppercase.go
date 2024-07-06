package main
import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	var inputString string
	 
	fmt.Println("Program to convert string to upper case: ")
	fmt.Println("Input a string: ")
	// takes input upto a new line
	reader := bufio.NewReader(os.Stdin)
	inputString, _ = reader.ReadString('\n')
	fmt.Println(strings.ToUpper(inputString))
}