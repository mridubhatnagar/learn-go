// solution to exercism problem on prime factors
package main
import "fmt"

func main() {
	// pass any number whose prime factors you want
	n := 60
	result := Factors(int64(n))
	fmt.Printf("Prime factors of %d are: %v\n", n, result)
}

func Factors(n int64) []int64 {
	var factors = []int64{}
    var divisor int64 = 2
    var remainder int64
    var quotient int64 
    for  n != 1 {
        remainder = int64(n)%int64(divisor)
        if remainder != 0 {
            divisor += 1
            continue
        } else {
            factors = append(factors, divisor)
        }
        quotient = n/divisor
        n = quotient
    }
    return factors
}