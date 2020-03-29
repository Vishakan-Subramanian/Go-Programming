package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//INTEST : https://www.codechef.com/problems/INTEST
func INTEST() {
	var n, count, i, k int

	count = 0

	fmt.Fscan(os.Stdin, &n)
	fmt.Fscan(os.Stdin, &k)

	reader := bufio.NewReader(os.Stdin)

	for i = 0; i < n; i++ {
		readValN, _ := reader.ReadString('\n')
		readVal := strings.TrimSuffix(readValN, "\n")

		t, _ := strconv.Atoi(readVal)

		if t%k == 0 {
			count++
		}
	}

	fmt.Println(count)
}
