package main 
import "fmt"

type DNA string


func main() {
	var d DNA
	d = "ATCG"
	fmt.Printf("%s\n", d.complement()) }

func (v DNA) complement() DNA {
	var s DNA = ""
	for i:=0;i<len(v);i++ {
		if v[i] == 'A' {
			s += "T"
		} else if v[i] == 'T' {
			s += "A"
		} else if v[i] == 'C' {
			s += "G"
		} else if v[i] == 'G' {
			s += "C"
		}
	}
	return s 
}
