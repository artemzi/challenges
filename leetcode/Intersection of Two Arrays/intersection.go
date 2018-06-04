package intersection

/**
Given two arrays, write a function to compute their intersection.

	Example:
		Given nums1 = [1, 2, 2, 1], nums2 = [2, 2], return [2].

	Note:
		Each element in the result must be unique.
		The result can be in any order.
*/

func intersection(nums1 []int, nums2 []int) []int {
	set := make([]int, 0)
	hash := make(map[int]bool)

	for i := range nums1 {
		hash[nums1[i]] = true
	}

	for i := range nums2 {
		if _, found := hash[nums2[i]]; found {
			set = append(set, nums2[i])
			delete(hash, nums2[i])
		}
	}

	return set
}
