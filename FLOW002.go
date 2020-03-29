package main

import (
	"fmt"
	"os"
)

//FLOW002 : https://www.codechef.com/problems/FLOW002
func FLOW002() {
	var dividend, divisor, tests int

	fmt.Fscan(os.Stdin, &tests)

	for i := 0; i < tests; i++ {
		fmt.Fscan(os.Stdin, &dividend)
		fmt.Fscan(os.Stdin, &divisor)
		fmt.Println(dividend % divisor)
	}
}
