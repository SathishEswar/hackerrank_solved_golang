package main

import (
    "fmt"
    "strings"
)

func main(){
    var n int 
    fmt.Scan(&n)
    var spaces string
    var hases string
    
    for i :=1 ; i<=n ;i++{
        spaces = strings.Repeat(" ", n - i)
        hases = strings.Repeat("#", i)
        fmt.Println( spaces + hases )

    }
    
}
