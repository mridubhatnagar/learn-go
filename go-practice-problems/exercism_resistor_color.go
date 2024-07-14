// solution to resistor color problem on exercism

package main 
import "fmt"

func main() {
	// initialization can be changed to any color.
	var color string = "orange"
	colorIndex := ColorCode(color)
	fmt.Printf("Color index for %s is %d\n", color, colorIndex)
}
// Colors returns the list of all colors.
func Colors() []string {
	return []string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
    var colors [] string
	colors = Colors()
    for index, value := range colors {
        if color == value {
            
            return index
        }
    }
    return -1
}
