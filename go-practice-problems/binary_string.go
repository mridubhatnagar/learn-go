package main

import (
	"fmt"
	"strconv"
	"math"
	"errors"
)

func main() {
	var inputNumber int = 13
	var binaryString string 
	var number int
	var err error  
	binaryString = intergerToBinary(inputNumber)
	fmt.Printf("Integer %d to binary %s\n", inputNumber, binaryString)
	number, err = binaryToInteger(binaryString)
	fmt.Printf("Binary %s to Integer %d. Errors: %v\n", binaryString, number, err)
	
}

func intergerToBinary(n int) string {
	var r int 
	var q int 
	var binaryStr string = ""
	for n > 0 {
		r = n%2 
		q = n/2 
		n = q 
		if r == 1 {
			binaryStr = "1" + binaryStr
		} else if r == 0 {
			binaryStr = "0" + binaryStr
		}
	}
	return binaryStr
}

func binaryToInteger(n string) (int, error) {
	var output int = 0
	var err error
	var result int  
	for i:=len(n)-1;i>=0;i-- {
		result, err = strconv.Atoi(string(n[i]))
		if err != nil {
			return 0, errors.New("Atoi returned error")  
		}
		output += int(math.Pow(float64(2), float64(len(n)-1-i))) * result 
	}
	return output, err 
}