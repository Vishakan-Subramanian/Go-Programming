package main

import (
	"fmt"
	"os"
	"sort"
)

//ZCO14003 : https://www.codechef.com/LRNDSA01/problems/ZCO14003
func ZCO14003() {
	var nCust, nBuyers int

	fmt.Fscan(os.Stdin, &nCust)
	nBuyers = nCust
	budget := make([]int, nCust)

	for i := 0; i < nCust; i++ {
		fmt.Fscan(os.Stdin, &budget[i])
	}
	sort.Ints(budget)
	for i := 0; i < nCust; i++ {
		budget[i] = budget[i] * nBuyers
		nBuyers--
	}
	sort.Ints(budget)
	fmt.Println(budget[nCust-1])
}
