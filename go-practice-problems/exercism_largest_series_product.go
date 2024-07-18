// solution to exercism largest series product
package main
import (
    "strconv"
    "fmt"
    "errors"
)

func main() {
	// digits input 
	digits := "63915"
	result, error := LargestSeriesProduct(digits, 3)
	fmt.Printf("Largest series product is %d. Error: %v\n", result, error)
}

func LargestSeriesProduct(digits string, span int) (int64, error) {
    var product int = 1
    var largestProduct int 
    var series string
    if len(digits) == span {
        for _, value := range digits {
            intVal, err := strconv.Atoi(string(value))
            if err != nil {
            	return 0, errors.New("strconv returned error for value")    
            }
            product *= intVal
        }
        largestProduct = product
        return int64(largestProduct), nil
    } else if span > len(digits) || digits == "" || span < 0  {
        var err error 
        err = errors.New("span must be smaller than string length")
        return 0, err
    } 
    
	for i:=0;i<=len(digits)-span;i++ {
        product = 1
        series = digits[i:i+span]
        for _, value := range series {
            // fmt.Println(string(value))
            intVal, err := strconv.Atoi(string(value))
            if err != nil {
            	return 0, errors.New("strconv returned error for value")    
            }
            product *= intVal
        }
        if product > largestProduct {
        	largestProduct = product    
        }
    }
    return int64(largestProduct), nil
}
