package main

import (
	"fmt"
	"sort"
)

func main(){
	var countmin,countmax int 	
	m:= make([]int, 5)
    for i :=0 ; i<5 ;i++{
		fmt.Scanf("%d", &m[i])
        
	}
	sort.Ints(m)

	for j :=0;j<5;j++{
		if j != (5-1){
			countmin = countmin+m[j]
		}
		if j!= 0 {
			countmax = countmax+m[j]
		}


	}
	fmt.Print(countmin,countmax)
    
}

