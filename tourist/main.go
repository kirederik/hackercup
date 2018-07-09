package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var testCases int
	if _, err := fmt.Scanf("%d\n", &testCases); err != nil {
		panic(err)
	}

	for i := 0; i < testCases; i++ {
		var (
			n, k int
			v    int64
		)
		if _, err := fmt.Scanf("%d %d %d\n", &n, &k, &v); err != nil {
			panic(err)
		}

		attractions := make([]string, n)
		for j := 0; j < n; j++ {
			var name string
			if _, err := fmt.Scanf("%s\n", &name); err != nil {
				panic(err)
			}
			attractions[j] = name
		}

		toVisit := attractionsToVisit(n, k, v, attractions)
		fmt.Printf("Case #%d: %s\n", i+1, toVisit)
	}
}

func attractionsToVisit(n, k int, v int64, attractions []string) string {
	var i int64
	i = int64(k) * (v - 1)
	indexesToVisit := make([]int, k)
	for j := 0; j < k; j++ {
		indexesToVisit[j] = (int(i) + j) % n
	}
	sort.Ints(indexesToVisit)

	var toVisit string
	for _, i := range indexesToVisit {
		toVisit = toVisit + attractions[i] + " "
	}

	return strings.TrimRight(toVisit, " ")
}
