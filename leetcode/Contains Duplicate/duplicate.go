package duplicate

func containsDuplicate(nums []int) bool {
	set := make(map[int]bool, len(nums))
	for _, n := range nums {
		if num, _ := set[n]; num == true {
			return true
		}
		set[n] = true
	}
	return false
}
