package main

import (
    
    "fmt"

)

var a = map[string]int{
    "a": 0,
    "b": 1,
    "c": 2,
    "d": 3,
    "e": 4,
    "f": 5,
    "g": 6,
    "h": 7,
    "i": 8,
    "j": 9,
    "k": 10,
    "l": 11,
    "m": 12,
    "n": 13,
    "o": 14,
    "p": 15,
    "q": 16,
    "r": 17,
    "s": 18,
    "t": 19,
    "u": 20,
    "v": 21,
    "w": 22,
    "x": 23,
    "y": 24,
    "z": 25,
}

func main() {
    va := make([]int, 26)
    for i := 0; i < 26; i++ {
        fmt.Scan(&va[i])
    }

    var s,k string
    fmt.Scan(&s)
    k = s
    var m int
    for l := range k {
        if z, ok := a[string(k[l])]; ok {
            if va[z] > m {
                m = va[z]
            }
        }
    }

    fmt.Println(m * 1 * len(s))
}

