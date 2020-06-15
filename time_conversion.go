package main

import (
    "fmt"
    "strconv"
)

func main() {
    var t string
    fmt.Scan(&t)
    h := t[0:2]
    m := t[3:5]
    s := t[6:8]
    a := string(t[8])
    ega := "A"
    egp := "P"

    H, _ := strconv.Atoi(h)

    switch {
    case H == 12 && a == ega:
        fmt.Printf("00:%v:%v", m, s)
    case a == ega:
        fmt.Printf("%v:%v:%v",h, m, s)
    case a ==egp && H==12:
        fmt.Printf("12:%v:%v", m, s)
    case a ==egp :
        val := H+12
        fmt.Printf("%v:%v:%v",val, m, s)

}
}
