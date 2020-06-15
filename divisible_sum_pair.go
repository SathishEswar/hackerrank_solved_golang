package main

import (
    "fmt"
    "sort"
)

func main() {
    var n, k, count int
    fmt.Scan(&n, &k)
    m := make([]int, n)

    for i := 0; i < n; i++ {
        fmt.Scan(&m[i])
    }
    sort.Ints(m)

    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            if (m[i]+m[j])%k == 0 {
                count++
            }
        }
       
    }
    fmt.Println(count)
}
