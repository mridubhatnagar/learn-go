package main
import "fmt"

func main() {
    var a = [2][3]int{{1,2,3},{4,5,6}}
    var b = [3][2]int{{7,8},{9,10},{11,12}}
    // var a = [1][3]int{{3,4,2}}
    // var b = [3][4]int{{13,9,7,15},{8,7,4,6},{6,4,0,3}}
    var sum int
    var c[2][2] int
    var rows1 int = 2
    var cols2 int = 2
    var rows2 int = 3
    for i:=0;i<rows1;i++ {
        for m:=0;m<cols2;m++ {
            sum=0
            for j:=0;j<rows2;j++ {
                sum += a[i][j] * b[j][m]
            }
            c[i][m] = sum
        }
    }
    for i:=0;i<rows1;i++ {
        for j:=0;j<cols2;j++ {
            fmt.Printf("%d ", c[i][j])
        }
        fmt.Printf("\n")
    }
}