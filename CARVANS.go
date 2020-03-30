package main

import (
	"fmt"
	"math"
	"os"
)

//CARVANS : https://www.codechef.com/LRNDSA01/problems/CARVANS
func CARVANS() {
	var testcases, noCars, maxSpeeds, curSpeed int

	fmt.Fscan(os.Stdin, &testcases)

	for i := 0; i < testcases; i++ {
		fmt.Fscan(os.Stdin, &noCars)
		maxSpeeds = 0
		speed := math.MaxInt32

		for j := 0; j < noCars; j++ {
			fmt.Fscan(os.Stdin, &curSpeed)
			if curSpeed <= speed {
				maxSpeeds++
				speed = curSpeed
			}
		}

		fmt.Println(maxSpeeds)
	}

}
