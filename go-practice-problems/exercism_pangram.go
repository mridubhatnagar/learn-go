package main 
import "fmt"

func main() {
	var input string = "abcdegfhijklmnopqrstuvwxyz"
	result:=IsPangram(input)
	fmt.Println(result)
}

func IsPangram(input string) bool {
	var isPangram bool = true 
	var ctr int 
	for i:='a';i<='z'; i++ {
		ctr = 0
		for _, value := range input {
			if i == value {
				ctr += 1
			}
		}
		if ctr == 0 {
			isPangram = false 
			break 
		}
	}
	return isPangram
}