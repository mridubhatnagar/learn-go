package main

import (
    "fmt"
    "math"
    "errors"
)

func main() {
	var grainCountOnSquare uint64 
	var number int = 12
	var err error 
	var totalGrainCount uint64 
	grainCountOnSquare, err = Square(number)
	fmt.Printf("grain count on square number %d is %d\n", number, grainCountOnSquare)
	fmt.Printf("Error grrr: %v\n", err)
	totalGrainCount = Total()
	fmt.Printf("Total Grains on board: %d\n", totalGrainCount)

}


func Square(number int) (uint64, error) {
    var grain uint64
    var err error = nil
    if number < 0 {
        err = errors.New("Number cannot be less than 0")
        return uint64(number), err
    } else if number == 0 {
        err = errors.New("Number cannot 0")
        return uint64(number), err
    } else if number > 64 {
        err = errors.New("Number cannot greater than 64. As board has only 64 squares.")
        return uint64(number), err
    }
	
    grain = uint64(math.Pow(float64(2), float64(number-1)))
    return grain, err
}

func Total() uint64 {
    var totalGrain uint64
    for i:=0;i<64;i++ {
        totalGrain += uint64(math.Pow(float64(2), float64(i)))
    }
	return totalGrain
}
