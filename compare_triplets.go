package main

import(
	"fmt"
)

func main(){
	var  a [3]int
	var  b [3]int
	var alice, bob int

	for i:=0;i<3;i++{
	   fmt.Scanf("%v",&a[i])
	}

	for i:=0;i<3;i++{
		fmt.Scanf("%v",&b[i])
	}

    for i := 0 ; i<3 ; i++{
		if a[i]>b[i]{
			alice++
		} else if a[i]<b[i] {
			bob++
		}
	}
	fmt.Println(alice,bob)
}
