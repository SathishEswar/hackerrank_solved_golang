package main

import(
    "fmt"
)


func main(){
    var y int
    fmt.Scan(&y)
    switch {
    case y >= 1700 && y < 1918 :
        if y%4 ==0{
            fmt.Printf("12.09.%d", y)
        }else{
            fmt.Printf("13.09.%d", y)
        }
    case y ==1918 :
        fmt.Printf("26.09.%d", y)
    case y>1918 && y<2701 :
        if y%400 ==0 || (y%4 ==0 && y%100 !=0) {
            fmt.Printf("12.09.%d", y)
        }else{
            fmt.Printf("13.09.%d", y)
        }

    }
}



