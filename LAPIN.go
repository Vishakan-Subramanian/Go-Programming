package main

import (
	"fmt"
	"os"
	s "strings"
)

//LAPIN : https://www.codechef.com/LRNDSA01/problems/LAPIN
func LAPIN() {
	var tests, wLen int
	var fHalf, sHalf, word string

	fmt.Fscan(os.Stdin, &tests)

	for i := 0; i < tests; i++ {
		fmt.Scan(&word)
		wLen = len(word)
		fHalf = word[0 : wLen/2+wLen%2]
		sHalf = word[wLen/2 : wLen]

		flag := false
		for _, j := range fHalf {
			c := string(j)
			if s.Count(fHalf, c) != s.Count(sHalf, c) {
				flag = true
				break
			}
		}
		if flag == true {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}
