package main

import (
	"fmt"
	"os"
)

//FLOW006 : https://www.codechef.com/problems/FLOW006
func FLOW006() {
	var x, tests, sum int

	fmt.Fscan(os.Stdin, &tests)

	for i := 0; i < tests; i++ {
		fmt.Fscan(os.Stdin, &x)
		sum = 0
		for true {
			if x <= 0 {
				break
			} else {
				sum += x % 10
				x /= 10
			}
		}
		fmt.Println(sum)
	}
}
