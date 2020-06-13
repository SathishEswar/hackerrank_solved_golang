package main

import (
    "fmt"
)

func main() {
    var sentence string
    var n, result int
    fmt.Scanln(&sentence)
    fmt.Scan(&n)

    for _, k := range sentence {
        if k == 'a' {
            result++
        }
    }

    result1 := (n / len(sentence)) * result

    for _, k := range sentence[:n%len(sentence)] {
        if k == 'a' {
            result1++
        }
    }

    fmt.Println(result1)
}
