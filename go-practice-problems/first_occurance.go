package main 
import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	var inputString string 
	var subString string 
	fmt.Println("Display index of first occurance of substring")
	fmt.Println("Enter main string: ")
	reader := bufio.NewReader(os.Stdin)
	inputString, _ = reader.ReadString('\n')
	fmt.Println("Enter substring: ")
	subString, _ = reader.ReadString('\n')
	subString = strings.TrimSuffix(subString, "\n")
	fmt.Println(strings.Index(inputString, subString))
}