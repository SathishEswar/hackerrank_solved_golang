
package main

import(
"fmt"

"strconv"
)

func main(){
var n,count int
var k int
var s string
fmt.Scan(&n)
for i:=0;i<n;i++{
    fmt.Scan(&k)
    count = 0
    s = strconv.Itoa(k)
   
    for j := 0;j<len(s);j++ {
        v, _ := strconv.Atoi(string(s[j]))
        if v!=0 && k%v==0{
            count++
        }
        }
        fmt.Println(count)       
    }
}
