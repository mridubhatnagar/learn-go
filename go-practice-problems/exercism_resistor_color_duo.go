// solution resistor color duo problem on exercism
package main 
import "fmt"

func main() {
	var colors = []string{"black", "brown"}
	result := Value(colors)
	fmt.Printf("Resistance band value for %v is %d", result)
}

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
    var bandValue int
    var possibleColors = [10]string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}
    for index, value := range colors {
        for index1, value1 := range possibleColors {
			// ten's place
            if value == value1 && index == 0 {
                bandValue += 10 * index1
				break
			// unit's place	
            } else if value == value1 && index == 1 {
                bandValue += 1 * index1
                break
            }
        }
    }
    return bandValue
}
