package main

import (
	"fmt"
	"os"
)

//FLOW004 : https://www.codechef.com/problems/FLOW004
func FLOW004() {
	var num, tests, sum int

	fmt.Fscan(os.Stdin, &tests)

	for i := 0; i < tests; i++ {
		sum = 0
		fmt.Fscan(os.Stdin, &num)
		sum = num % 10
		for true {
			if num/10 > 0 {
				num /= 10
			} else {
				break
			}
		}
		sum += num
		fmt.Println(sum)
	}
}
