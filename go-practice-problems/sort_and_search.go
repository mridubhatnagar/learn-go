// Write the algorithm for binary search to work on an array of structures.
// you would load the CSV file to an array of structures
// Sort on rollno
// and then use the sorted array for the binary search

// pos=binarySearch(sortedArray, search_rollnumber)
// If it is found, return the position, otherwise return -1
package main 
import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
)

type student_record struct {
	rollno	int
	name string 
	enrollement_year int 
	grade string 
	cgpa float64 
}

func main() {
	readFile, err := os.Open("student_records_1.csv")
	var lineNumber int = 0 
	var records []student_record
	var sortedArray []student_record
	var pos int 
	var searchRollno int = 12
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
        records = createStructure(fileScanner.Text(), lineNumber, records)
		lineNumber += 1
    }
	readFile.Close()
	sortedArray = sortRecords(records)
	pos = binarySearch(sortedArray, searchRollno)
	if pos != -1 {
		fmt.Printf("Roll no. %d found at index: %d\n", searchRollno, pos)
	} else {
		fmt.Printf("Roll no. %d not found\n")
	}
	

}

func createStructure(line string, lineNumber int, records []student_record) []student_record {
	var r student_record
	var lineArray = strings.Split(line, ",")
	if lineNumber >= 1 {
		r.rollno, _ = strconv.Atoi(lineArray[0])
		r.name  = lineArray[1]
		r.enrollement_year, _  = strconv.Atoi(lineArray[2])
		r.grade = lineArray[3]
		r.cgpa, _ = strconv.ParseFloat(lineArray[4], 64)
		records = append(records, r)
	}
	return records
}

func sortRecords(records []student_record) []student_record {
	var minIndex int 
	var temp student_record 
	for i:=0;i<=len(records)-2;i++ {
		minIndex = i 
		for j:=i+1;j<=len(records)-1;j++ {
			if records[j].rollno < records[minIndex].rollno {
				minIndex = j
			}
		}
		temp = records[minIndex]
		records[minIndex] = records[i]
		records[i] = temp 
	}
	return records
}

func binarySearch(sortedArray []student_record, rollno int) int{
	var low int = 0
	var high int = len(sortedArray)-1
	var mid int 
	for low <= high {
		mid = (low+high)/2
		if rollno == sortedArray[mid].rollno {
			return mid 
		} else if rollno < sortedArray[mid].rollno {
			high = mid - 1
		} else if rollno > sortedArray[mid].rollno {
			low = mid + 1
		} 
	}
	return -1
}
