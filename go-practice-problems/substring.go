package main 
import (
	"fmt"
	"bufio"
	"strings"
	"os"
)

func main() {
	var inputString string 
	var subString string 
	fmt.Println("Find substring exists or not: ")
	fmt.Println("Input the main string: ")
	reader := bufio.NewReader(os.Stdin)
	inputString, _ = reader.ReadString('\n')
	fmt.Println("Input string to be searched: ")
	subString, _  = reader.ReadString('\n')
	// doesn't find word which have space in beginning and end.
	// TrimSuffix removes the end string if it matches. Here \n.
	fmt.Println(strings.Contains(inputString, strings.TrimSuffix(subString, "\n")))
	fmt.Println(strings.TrimSuffix(inputString, "ran\n"))
}