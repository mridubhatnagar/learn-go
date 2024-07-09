package main
import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var inputString string 
	var startIndex int 
	var size int 
	var newString string 
	fmt.Println("Display part of string based on start index and size: ")
	fmt.Println("Input main string: ")
	reader := bufio.NewReader(os.Stdin)
	inputString, _ = reader.ReadString('\n')
	fmt.Println("Enter start index: ")
	fmt.Scan(&startIndex)
	fmt.Println("Enter the size of string you want to display")
	fmt.Scan(&size)
	newString = inputString[startIndex: size]
	fmt.Printf("%s\n", newString)

}