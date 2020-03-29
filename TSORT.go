package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

//TSORT : https://www.codechef.com/problems/TSORT
func TSORT() {
	var len int
	var list []int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	fmt.Fscan(os.Stdin, &len)

	for i := 0; i < len; i++ {
		readValN, _ := reader.ReadString('\n')
		readVal := strings.TrimSuffix(readValN, "\n")
		val, _ := strconv.Atoi(readVal)
		list = append(list, val)
	}

	sort.Ints(list)
	for i := 0; i < len; i++ {
		strVal := strconv.Itoa(list[i])
		writer.WriteString(strVal + "\n")
	}
	writer.Flush()
}
