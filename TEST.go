package main

import (
	"fmt"
	"os"
)

//TEST : https://www.codechef.com/LRNDSA01/problems/TEST
func TEST() {
	var x int

	x = 0
	for true {
		fmt.Fscan(os.Stdin, &x)
		if x == 42 {
			break
		} else {
			fmt.Println(x)
		}
	}

}
