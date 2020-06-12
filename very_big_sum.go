package main
import "fmt"

func main() {
    var n int
    fmt.Scanf("%v", &n)
    
    var res uint64
    for i := 0; i < n; i++ {
        var num uint64
        fmt.Scanf("%v", &num)
        res +=num
    }
    
    fmt.Println(res)
}
