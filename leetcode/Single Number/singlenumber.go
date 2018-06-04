// https://leetcode.com/explore/learn/card/hash-table/183/combination-with-other-algorithms/1176/
package singlenumber

func singleNumber(nums []int) int {
	set := make(map[int]int, len(nums))
	for _, n := range nums {
		set[n]++
	}

	for k, val := range set {
		if val == 1 {
			return k
		}
	}
	return -1
}
