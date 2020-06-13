package main

import "fmt"

func main() {
    var n,k,res int
    fmt.Scanf("%d", &n)
    fmt.Scanf("%d", &k)
     res = 0
    var c = make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&c[i])
        res = res + c[i]
    }

    var b int
    fmt.Scanf("%d", &b)
    if b == res/2 {
        fmt.Print(c[k] / 2)
    } else {
        fmt.Print("Bon Appetit")
    }
    }
