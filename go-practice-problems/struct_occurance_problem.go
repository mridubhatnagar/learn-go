package main 
import "fmt"

func main() {
	var inputString string = "helloworld"
	var ctr int 
	type letterCount struct {
		letter string
		count  int 
	}
	// var x letterCount 
	var result[]letterCount

	for i:='a'; i <= 'z'; i++ {
		ctr = 0
		for _, value := range inputString {
			if i == value {
				ctr += 1
			}
		}
		if ctr >= 1 {
			result = append(result, letterCount {
				letter: string(i),
				count: ctr,
			})
		}

	}
	fmt.Printf("%v\n", result)
}

