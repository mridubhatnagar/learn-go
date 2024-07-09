// Q. Write a program to display the position of all occurrences of a substring in a string


package main
import (
    "fmt"
    "strings"
)

func main() {
  var ctr int
  var index int
  var remaining_string string
  var prevIndex int = -1
  inputString := "A cat ran after a dog but the dog ran faster"
  subString := "a"
  ctr = 0
  for {
      index = strings.Index(inputString, subString)
      fmt.Printf("Index: %d\n", index)
      remaining_string = inputString[index+1:]
     
      fmt.Printf("Remaining string: %s\n", remaining_string)
      inputString = remaining_string
      if index == -1 {
          break
      }
      if ctr >= 1 {
            index = prevIndex + index+1    
      }
      ctr += 1
      prevIndex = index
      fmt.Printf("Index of %d occurrence: %d \n", ctr, index)
  }
}