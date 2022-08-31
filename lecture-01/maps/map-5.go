package maps

import (
	"fmt"
	"sort"
)

// You have map[int1]map[int2]string, you need to print it as sorted data: descending sorting by fee first and ascending by nonce last.

func Map_5(m map[int]map[int]string) map[int]map[int]string {
	fmt.Println("You have map[int1]map[int2]string, you need to print it as sorted data: descending sorting by int1 first and ascending by int2 last.")
	fmt.Println("Map_5 - Map:", m)

	keys1 := make([]int, 0, len(m))
	for k1 := range m {
		keys1 = append(keys1, k1)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(keys1)))

	result := make(map[int]map[int]string, len(keys1))

	for _, k1 := range keys1 {
		keys2 := make([]int, 0, len(m[k1]))

		for k2 := range m[k1] {
			keys2 = append(keys2, k2)
		}

		result[k1] = make(map[int]string, len(keys2))

		sort.Ints(keys2)

		for _, k2 := range keys2 {
			fmt.Println("Map_5 : ", k1, k2, m[k1][k2])
			result[k1][k2] = m[k1][k2]
		}
	}

	return result
}
