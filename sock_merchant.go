package main
import "fmt"

var sf = fmt.Scanf
var p = fmt.Println
var pf = fmt.Print

func main() {
    
    var (
        n, key, s int
        m map[int] int    
    )
    
    if _,err := sf("%d", &n); err != nil {
        panic(err)
    }
    
    m = make(map[int]int)
    
    for i:=0; i < n; i++ {
        if _,err := sf("%d", &key); err != nil {
            panic(err)
        }
        if _,ok := m[key]; ok {
            m[key]++
        } else {
            m[key] = 1    
        }        
    }
   /* for k, v := range m {
        p("Key:", k, " Value:",v)
    } */
    
    s = 0
    for _, v := range m {
        s += v/2
    }
    p(s)
}
