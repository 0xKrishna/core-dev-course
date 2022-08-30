package maps

import "fmt"

// There is a very large array or slice of integers, it must be said which numbers are mentioned in it at least once.

func Map_2(s []int) {
	fmt.Println("There is a very large array or slice of integers, it must be said which numbers are mentioned in it at least once.")
	fmt.Println("Map_2 - Slice:", s)

	numbers := make(map[int]struct{})

	for _, number := range s {
		numbers[number] = struct{}{}
	}

	for number := range numbers {
		fmt.Println("Map_2 - Number:", number)
	}
}
