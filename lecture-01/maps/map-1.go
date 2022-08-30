package maps

import "fmt"

// There is a text, you need to count how many times each word occurs.

func Map_1(s string) {
	fmt.Println("There is a text, you need to count how many times each word occurs.")
	fmt.Println("Map_1 - Text:", s)

	words := make(map[string]int)

	for _, word := range s {
		words[string(word)]++
	}

	for word, count := range words {
		fmt.Println("Map_1 - Word:", word, "Count:", count)
	}
}
