// Use double linked list to read data from CSV. 
// Store struct in data part of node in sorted order of rollno.  

package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type srecord struct {
	rollno	int
	name string 
	enrollement_year int 
	grade string 
	cgpa float64 
}

type snode struct {
	data srecord
	prev *snode
	next *snode 
}


func main() {
	// step 1. Open file
	readFile, err := os.Open("student_records_1.csv")
	var lineNumber int = 0 
	var recordStruct srecord
	var snoderecord *snode
	var start *snode 
	var end *snode 
	var result *snode
	var searchKey int = 10
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
        recordStruct = createStudentStruct(fileScanner.Text(), lineNumber)
		snoderecord = SNode(recordStruct)
		start, end = insertSortedStudentStruct(start, end, snoderecord)
		lineNumber += 1
    }
	readFile.Close()
	traverseLforward(start)
	traverseLBackward(end)
	result = searchStudentRecord(start, searchKey)
	if result != nil {
		fmt.Printf("Roll no. %d details are: ", result.data.rollno)
		fmt.Printf("Name: %s, Enrollment Year: %d, CGPA: %.2f, Grade: %s\n", result.data.name, result.data.enrollement_year, result.data.cgpa, result.data.grade)
	} else {
		fmt.Printf("Roll no. %d not found.", result.data.rollno)
	}

}

func SNode(data srecord) *snode{
	n := &snode{data, nil,nil}
	return n 
}

func createStudentStruct(line string, lineNumber int) srecord {
	var r srecord
	var lineArray = strings.Split(line, ",")
	if lineNumber >= 1 {
		r.rollno, _ = strconv.Atoi(lineArray[0])
		r.name  = lineArray[1]
		r.enrollement_year, _  = strconv.Atoi(lineArray[2])
		r.grade = lineArray[3]
		r.cgpa, _ = strconv.ParseFloat(lineArray[4], 64)
	}
	return r
}

func traverseLforward(start *snode) {
	var p *snode 
	for p=start;p!=nil;p=p.next{
		displayNodeRecord(p)
	}
	fmt.Printf("\n")
}

func traverseLBackward(end *snode) {
	var p *snode 
	for p=end;p!=nil;p=p.prev {
		displayNodeRecord(p)
	}
	fmt.Printf("\n")
}

func displayNodeRecord(p *snode) {
	fmt.Printf("%d %s\n", p.data.rollno, p.data.name)
}



func insertSortedStudentStruct(start *snode, end *snode, n*snode) (*snode, *snode) {
	var p *snode 
	var t *snode 
	if start == nil {
		n.next = nil
		n.prev = nil
		start = n 
		end = n
	} else {
		for p=start; p!=nil;p=p.next {
			if p.data.rollno > n.data.rollno {
				if p == start {
					n.next = start
					start.prev = n
					n.prev = nil  
					start = n 
					break 
			 	} else {
					t = p.prev
					t.next = n
					n.next = p
					n.prev = t
					p.prev = n
					break
		   		}
			}

		}
		if p == nil {
			end.next = n
			n.prev = end
			n.next = nil    
			end = n
		}
	}
	return start, end 
}


func searchStudentRecord(start *snode, data int)  *snode {
	var p *snode 
	for p=start;p!=nil;p=p.next {
		if p.data.rollno == data {
			return p 
		}
	}
	return nil
}