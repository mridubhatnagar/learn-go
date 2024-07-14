// solution to exercism problem on collatz conjuncture
package main
import (
    "errors"
    "fmt"
)

func main() {
	// for testing keep changing the value of number passed.
	number, error:= CollatzConjecture(12)
	fmt.Printf("Step count: %d, error: %v\n", number, error)
}


func CollatzConjecture(n int) (int, error) {
    var ctr int = 0
    var err error 
    if n == 0 {
        err = errors.New("n is 0")
        return n, err
    } else if n < 0 {
        err = errors.New("n cannot be negative")
        return n, err
    } else {
        err = nil
    }
	for  n != 1 {
        if n%2 == 0 {
            n = n/2
        } else {
            n = 3*n + 1
        }
        ctr += 1
    }
    return ctr, err
}
