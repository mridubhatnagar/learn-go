// Question 01
// =========
// Load a CSV file into a list of structures
// Make a CSV file through spreadsheet and save as CSV
// There should be a heading field ( lower case, single word )
// We can use :
// rollno,name,enrollment_year,grade,cgpa

// - enrollment_year will be numeric like 2001, 2005 etc
// - cgpa will be numeric with decimal
// Put some data ( around 10 records )
// Assume that the data will be valid

// So define a struct that corresponds to the header of the csv, for name and data type
// After the csv is successfully loaded into a list of such structures:
// a. Display all the records and complete data ( for verification )
// b. Display all records ( rollno, name, cgpa ) for cgpa > someValue
// c. Display all records ( rollno, name ) for year that is equal to certain year

// sort array of structure based on cgpa
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type record struct {
	rollno	int
	name string 
	enrollement_year int 
	grade string 
	cgpa float64 
}


func main() {
	// step 1. Open file
	readFile, err := os.Open("student_records.csv")
	var lineNumber int = 0 
	var searchKey = make(map[int]int)
	var records []record
	var sortedArray []record
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
        records = createStruct(fileScanner.Text(), lineNumber, records)
		lineNumber += 1
    }
	readFile.Close()
	fmt.Println(records)
	filterRecordsByCgpa(records)
	filterRecordsByEnrollmentYear(records, 2013)
	searchKey = searchMap(records, searchKey)
	fmt.Println(searchKey)
	sortedArray = sortArrayByCgpa(records)
	fmt.Println(sortedArray)
}

func createStruct(line string, lineNumber int, records []record) []record {
	var r record
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

func filterRecordsByCgpa(records []record) {
	for i:=0;i<len(records); i++ {
		if records[i].cgpa > 9.0 {
			fmt.Printf("Roll no: %d, Name: %s, CGPA: %.2f\n", records[i].rollno, records[i].name, records[i].cgpa)
		}
	}

}

func filterRecordsByEnrollmentYear(records []record, enrollment_year int) {
	for i:=0;i<len(records); i++ {
		if records[i].enrollement_year == enrollment_year {
			fmt.Printf("Roll no: %d, Name: %s\n", records[i].rollno, records[i].name)
		}
	}
}

// solution to problem 3. Create map with {rollno. : arrayIndex}
func searchMap(records []record, searchKey map[int]int) map[int]int{
	for i:=0;i<len(records);i++ {
			searchKey[records[i].rollno] = i
	}
	return searchKey
}


// sort array of struct based on cgpa
func sortArrayByCgpa(inputArray []record) []record {
	var temp record 
	var minimum record
	for i:=0;i<=len(inputArray)-1;i++ {
		minimum = inputArray[i]
		for j:=i+1;j<=len(inputArray)-2;j++{
			if inputArray[j].cgpa < minimum.cgpa {
				temp = inputArray[j]
				inputArray[j] = minimum
				inputArray[i] = temp
				minimum = inputArray[i]
			}
		}
	}
	return inputArray
}