package main

import (
	"fmt"
	"os"
)

//FLOW007 : https://www.codechef.com/problems/FLOW007
func FLOW007() {
	var num, tests, reverse, rem int

	fmt.Fscan(os.Stdin, &tests)

	for i := 0; i < tests; i++ {
		reverse = 0
		fmt.Fscan(os.Stdin, &num)
		for num > 0 {
			rem = num % 10
			reverse = reverse*10 + rem
			num /= 10
		}
		fmt.Println(reverse)
	}
}
