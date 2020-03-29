package main

import (
	"fmt"
	"os"
)

//LUCKFOUR : https://www.codechef.com/problems/LUCKFOUR
func LUCKFOUR() {
	var num, tests, count, rem int

	fmt.Fscan(os.Stdin, &tests)

	for i := 0; i < tests; i++ {
		count = 0
		fmt.Fscan(os.Stdin, &num)
		for num > 0 {
			rem = num % 10
			if rem == 4 {
				count++
			}
			num /= 10
		}
		fmt.Println(count)
	}
}
