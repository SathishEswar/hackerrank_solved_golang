package main

import(
"fmt"
)

func main(){

    var n,k,thatdayad,cumulative int
	k=5
	cumulative =2
    fmt.Scan(&n)
	if n==1{
		fmt.Println("2")
	}else{
    for i:=2 ;i<n+1 ;i++{
            thatdayad = (k/2)*3
            k = thatdayad
            cumulative = cumulative+k/2
        }
		fmt.Println(cumulative) 
    }
       
}
