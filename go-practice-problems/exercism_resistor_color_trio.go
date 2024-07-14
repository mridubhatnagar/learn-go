// solution to exercism problem resistor color trio
package main
import (
    "strconv"
    "math"
)
// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
    var bandValue int = 0
    colors = colors[:3]
	var possibleColors = [10]string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}
    for index, value := range colors {
        for index1, value1 := range possibleColors {
            if value == value1 && index == 0 {
                bandValue += 10 * index1
                break
            } else if value == value1 && index == 1 {
                bandValue += 1 * index1
                break
            } else if value == value1 && index == 2 {
                bandValue *= int(math.Pow(float64(10), float64(index1)))
                break
            } 
        }
    }
    if bandValue > 100000000 {
    	return strconv.Itoa(bandValue/1000000000) + " gigaohms"
    } else if bandValue > 1000000 && bandValue < 1000000000 {
        return strconv.Itoa(bandValue/1000000) + " megaohms"
    } else if bandValue > 1000 && bandValue < 1000000 {
        return strconv.Itoa(bandValue/1000) + " kiloohms"
    } else {
        return strconv.Itoa(bandValue) + " ohms"
    }
}
