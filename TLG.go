package main

import (
	"fmt"
	"os"
)

//TLG : https://www.codechef.com/problems/TLG
func TLG() {
	var p1Score, p2Score, lead, currentWinner, tests, p1Total, p2Total int

	currentWinner, lead = 1, 0
	p1Total, p2Total = 0, 0
	fmt.Fscan(os.Stdin, &tests)

	for i := 0; i < tests; i++ {
		fmt.Fscan(os.Stdin, &p1Score)
		fmt.Fscan(os.Stdin, &p2Score)
		p1Total += p1Score
		p2Total += p2Score

		if p1Total-p2Total > lead {
			lead = p1Total - p2Total
			currentWinner = 1
		} else if p2Total-p1Total > lead {
			lead = p2Total - p1Total
			currentWinner = 2
		} else {
			continue
		}
	}
	fmt.Printf("%d %d", currentWinner, lead)
}
