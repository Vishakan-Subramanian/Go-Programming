package main

import (
	"fmt"
	"math/big"
	"os"
)

//FCTRL2 : https://www.codechef.com/problems/FCTRL2
func FCTRL2() {
	var tests, x int

	fmt.Fscan(os.Stdin, &tests)

	for i := 0; i < tests; i++ {
		fmt.Fscan(os.Stdin, &x)
		computeFactorial(x)

	}
}

func computeFactorial(x int) {
	var fact big.Int
	fmt.Println(fact.MulRange(1, int64(x)))
}
