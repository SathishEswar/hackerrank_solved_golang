package main
import(
	"fmt"
)

func main(){
	var m map[int]int
	m = make(map[int]int)
	fmt.Println(m)
	var x,n,res int
	res = 1
	
	fmt.Scan(&n)
	for i :=0 ; i < n; i++ {
		fmt.Scanf("%d", &x)
		m[x] = m[x]+1
		if m[x] > m[res]{
			res = x
		}else if m[x] == m[res] && m<res{
			rex = x
		}
	}

	fmt.Println(res)


}
