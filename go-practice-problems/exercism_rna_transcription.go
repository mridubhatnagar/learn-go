// solution to exercism problem on RNA transcription

package main 
import "fmt"

func main() {
	var dna string = "ACGGGTAAGG" 
	result := ToRNA(dna)
	fmt.Printf("Corresponding to DNA %s, RNA is %s\n", dna, result)
}

func ToRNA(dna string) string {
	var rna string 
	for _, value := range dna {
		if value == 'G' {
			rna += "C"
		} else if value == 'C'{
			rna += "G"
		} else if value == 'T' {
			rna += "A"
		} else if value == 'A' {
			rna += "U"
		}
	}
	return rna
}

