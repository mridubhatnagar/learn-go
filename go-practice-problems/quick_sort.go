package main 
import "fmt"


func main() {
	var inputArray = []int{10,80,30,90,40,50,70}
	var low int  = 0
	var high int = len(inputArray)-1
	result := quickSort(inputArray, low, high)
	fmt.Printf("Sorted Array: %v\n", result)
}

func quickSort(inputArray []int, low int, high int) []int{
	var pivotLocation int 

	if low < high {
		pivotLocation = divide(inputArray, low, high)
		inputArray = quickSort(inputArray, low, pivotLocation)
		inputArray = quickSort(inputArray, pivotLocation+1, high)
	} 
	return inputArray
}

func divide(inputArray []int, low int, high int) int{
	var mid = (low+high)/2 
	var lastsmall int 
	var temp int 
	var pivot int 
	// assume mid is pivot. Swap a[low] with a[mid]
	temp = inputArray[mid]
	inputArray[mid] = inputArray[low]
	inputArray[low] = temp
	pivot = inputArray[low]
	// lastsmall and low point to same element.
	lastsmall = low 
	for i:=low+1;i<=high; i++{
		// pivot is leftmost element. 
		if inputArray[i] < pivot {
			lastsmall++
			temp = inputArray[i]
			inputArray[i] = inputArray[lastsmall]
			inputArray[lastsmall] = temp
		}
	}
	temp = inputArray[low]
	inputArray[low] = inputArray[lastsmall]
	inputArray[lastsmall] = temp
	return lastsmall
}