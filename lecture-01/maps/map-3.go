package maps

import "fmt"

// There are two large arrays of numbers, you need to find which numbers are mentioned in both.

func Map_3(s1 []int, s2 []int) []int {
	fmt.Println("There are two large arrays of numbers, you need to find which numbers are mentioned in both.")
	fmt.Println("Map_3 - Slice 1:", s1)
	fmt.Println("Map_3 - Slice 2:", s2)

	numbers := make(map[int]struct{})

	for _, number := range s1 {
		numbers[number] = struct{}{}
	}

	result := make([]int, 0, len(s1))

	for _, number := range s2 {
		if _, ok := numbers[number]; ok {
			fmt.Println("Map_3 - Number:", number)
			result = append(result, number)
		}
	}

	return result
}
