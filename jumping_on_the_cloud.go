package main
import "fmt"

func main() {
    //Enter your code here. Read input from STDIN. Print output to STDOUT
    var n int;
    fmt.Scan(&n);
    cs := make([]int, n)
    for i := range cs {
        fmt.Scan(&cs[i])
    }
    var m int
    var i int
    for i<n-1 {
        if i+2 > n-1 || cs[i+2] == 1 {
            i++
            m++
        } else {
            i += 2
            m++
        }
        // fmt.Println(i)
    }
    fmt.Println(m)
}
