package main

import (
	"fmt"
	"os"
)

//FLOW001 : https://www.codechef.com/problems/FLOW001
func FLOW001() {
	var x, y, tests int

	fmt.Fscan(os.Stdin, &tests)

	for i := 0; i < tests; i++ {
		fmt.Fscan(os.Stdin, &x)
		fmt.Fscan(os.Stdin, &y)
		fmt.Println(x + y)
	}
}
