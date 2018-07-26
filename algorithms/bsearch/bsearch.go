package bsearch

// Bsearch implements binary search algorithm
func Bsearch(haystack []int, needle int) int {
	var fst int              // first index
	lst := len(haystack) - 1 // last index

	for fst <= lst {

		m := fst + (lst-fst)/2 // get mid element

		if haystack[m] < needle {
			fst = m + 1
		} else {
			lst = m - 1
		}

	}

	if fst == len(haystack) || haystack[fst] != needle {
		return -1
	} else {
		return fst
	}
}
