package slices

import "fmt"

// Sort slice in the given order: direct, reverse, lexicographic (you'd need at least []int and []string).

func Slice_15(s []int, a []string) {
	fmt.Println("Sort slice in the given order: direct, reverse, lexicographic (you'd need at least []int and []string).")
	fmt.Println("Slice_15 - Slice int:", s)
	fmt.Println("Slice_15 - Slice string:", a)

	s = sortSliceDirect(s)
	fmt.Println("Slice_15 - Direct:", s)

	s = sortSliceReverse(s)
	fmt.Println("Slice_15 - Reverse:", s)

	a = sortSliceLexicographic(a)
	fmt.Println("Slice_15 - Lexicographic:", a)
}

func sortSliceDirect(s []int) []int {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}

	return s
}

func sortSliceReverse(s []int) []int {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] < s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}

	return s
}

func sortSliceLexicographic(s []string) []string {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}

	return s
}
