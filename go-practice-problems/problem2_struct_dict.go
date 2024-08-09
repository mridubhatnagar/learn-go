// Question 02
// =========
// Map a key field to a structure { key: structure }
// Using the same CSV file, modify the structure to exclude rollno from it
// Now load the CSV file into a mapping ( dictionary like structure ) such that we have:
//   { rollno : structure }
// So we will have as many keys in this dictionary as the number of rollnos and the values will be the corresponding data of the student

// Using this dictionary/mapping, display the same outputs as done in the question 01
package main 
import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"unsafe"
)

type recordDict struct {
	name string 
	enrollment_year int 
	cgpa float64
	grade string 
}

func main() {
	// make initializes map
	var records = make(map[int]recordDict)
	var rowNumber int = 0
	readFile, err := os.Open("student_records.csv")
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
        records = createStructMap(fileScanner.Text(), rowNumber, records)
		rowNumber += 1
    }
	fmt.Println(records)
	fmt.Println(unsafe.Sizeof(records))
}

func createStructMap(row string, rowNumber int, records map[int]recordDict) map[int]recordDict {
	var r recordDict
	var rowArray = strings.Split(row, ",") 
	if rowNumber >= 1 {
		r.name = rowArray[1]
		r.enrollment_year, _ = strconv.Atoi(rowArray[2])
		r.cgpa, _ = strconv.ParseFloat(rowArray[4], 64)
		r.grade = rowArray[3]
		rollno, _ := strconv.Atoi(rowArray[0])
		records[rollno] = r
	}
	return records
} 