// solution to exercism problem on acronym
package main

import (
    "fmt"
    "strings"
)

func main() {
	// give any string as input
	s := "a cat ran"
	result := Abbreviate(s)
	fmt.Printf("%s\n", result)
}
func Abbreviate(s string) string {
    var words = []string{}
    var acronym string
    // replaces hyphens with whitespaces
    s=strings.Replace(s, "-", " ", len(s))
    // removes all whitespaces from the string
    words = strings.Fields(s)
    for _, value := range words {
        // if a word has underscore as prefix. Removes it. 
        value = strings.TrimPrefix(value, "_")
        acronym += strings.ToUpper(string(value[0]))
    }
    // fmt.Printf("%s", acronym)
	return acronym
}
