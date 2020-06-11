package main

import "fmt"

func main() {
	var firststep, check, n, result int
	var steps string

	fmt.Scan(&n)
	fmt.Scan(&steps)

	if steps[0:1] == "U" {
		firststep = 1
	} else {
		firststep = -1
	}

	for i, j := 1, 2; i < len(steps); i, j = i+1, j+1 {
		check = firststep
		if steps[i:j] == "U" {
			firststep += 1
		} else {
			firststep -= 1
		}

		if firststep == 0 && check == -1 {
			result++
		}

	}
	fmt.Println(result)
}
