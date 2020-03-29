package main

import (
	"fmt"
	"os"
)

// ATM : https://www.codechef.com/problems/HS08TEST
func ATM() {
	var wdAmt int64
	var initBal, curBal float64

	fmt.Fscan(os.Stdin, &wdAmt)
	fmt.Fscan(os.Stdin, &initBal)

	curBal = initBal

	if wdAmt%5 == 0 {
		var wdTotal = float64(wdAmt) + 0.5

		if wdTotal <= initBal {
			curBal = initBal - wdTotal
		}
	}

	fmt.Printf("%.2f", curBal)
}
