package main

import (
	"fmt"
	"modernc.org/mathutil"
	"sort"
	"strings"
)

func main() {
	permutation("a")
	permutation("ab")
	permutation("abc")
	permutation("aabb")
}

func permutation(str string) sort.StringSlice {
	fmt.Printf("input '%s':\n", str)
	data := sort.StringSlice(strings.Split(str, ""))

	var result []string

	i := 0
	for mathutil.PermutationFirst(data); ; i++ {
		result = append(result, strings.Join(data, ""))

		if !mathutil.PermutationNext(data) {
			break
		}
	}

	fmt.Printf("function return '%s'\n", result)
	return result
}
