package main
import(
    "fmt"
    "math"
)
func main(){
var a,b,c float64
var n int
fmt.Scan(&n)

for i:=0 ;i<n;i++{
 fmt.Scan(&a,&b,&c)
 x2:= math.Abs(c-b)
 x1 := math.Abs(c-a) 
 if x1<x2{
     fmt.Println("Cat A")
 }else if x2<x1{
     fmt.Println("Cat B")
 }else if x1==x2{
     fmt.Println("Mouse C")
 }
 
}
}
