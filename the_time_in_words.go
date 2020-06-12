package main

import (
    "fmt"
)

func main() {

    var a map[int]string
    a = map[int]string{
        1:  "one",
        2:  "two",
        3:  "three",
        4:  "four",
        5:  "five",
        6:  "six",
        7:  "seven",
        8:  "eight",
        9:  "nine",
        10: "ten",
        11: "eleven",
        12: "twelve",
        13: "thirteen",
        14: "fourteen",
        15: "quarter",
        16: "sixteen",
        17: "seventeen",
        18: "eighteen",
        19: "nineteen",
        20: "twenty",
        21: "twenty one",
        22: "twenty two",
        23: "twenty three",
        24: "twenty four",
        25: "twenty five",
        26: "twenty six",
        27: "twenty seven",
        28: "twenty eight",
        29: "twenty nine",
    }

    var x, m int
    fmt.Scan(&x)
    fmt.Scan(&m)

    switch {
    case m == 0:
        fmt.Println(a[x], "o' clock")
    case m == 1:
        fmt.Println("one minute past",a[x])    
    case m == 15:
        fmt.Println("quarter past", a[x])
    case m == 30:
        fmt.Println("half past", a[x])
    case m > 0 && m <= 30:
        fmt.Println(a[m], "minutes past", a[x])
    case m == 45:
        fmt.Println("quarter to", a[x+1])
    case m > 30 && m < 59:
        fmt.Println(a[60-m], "minutes to", a[x+1])
    case m == 59:
        fmt.Println("one minute to", a[x+1])
    }
}
