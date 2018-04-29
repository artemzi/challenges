// Package friendsages solve contest problem
// https://leetcode.com/contest/weekly-contest-82/problems/friends-of-appropriate-ages/
package friendsages

import "sort"

// Some people will make friend requests. The list of their ages is given and ages[i] is the age
// of the ith person.

// Person A will NOT friend request person B (B != A) if any of the following conditions are true:

// age[B] <= 0.5 * age[A] + 7
// age[B] > age[A]
// age[B] > 100 && age[A] < 100
// Otherwise, A will friend request B.

// Note that if A requests B, B does not necessarily request A.
// Also, people will not friend request themselves.

// How many total friend requests are made?

// Notes:

// 1 <= ages.length <= 20000.
// 1 <= ages[i] <= 120.

func numFriendRequests(ages []int) int {
	// 83 / 83 test cases passed.
	// Status: Accepted
	// Runtime: 80 ms
	var req int
	data := ages[:]
	sort.Ints(data)

	for i, v := range data {
		if 0.5*float64(v)+7 > float64(v) {
			continue
		}
		s := BisectRight(data, int(0.5*float64(v)+7))
		t := BisectRight(data, v)
		if s <= i && i <= t {
			req += (t - s - 1)
		} else {
			req += (t - s)
		}
	}
	return req
}

// BisectRight is equivalent of Python's bisect_right()
func BisectRight(a []int, v int) int {
	return bisectRightRange(a, v, 0, len(a))
}

func bisectRightRange(a []int, v int, lo, hi int) int {
	s := a[lo:hi]
	return sort.Search(len(s), func(i int) bool {
		return s[i] > v
	})
}

// +++++++++++++++++ Research: ++++++++++

// BisectLeft is equivalent of Python's bisect_left()
// func BisectLeft(a []int, v int) int {
// 	return bisectLeftRange(a, v, 0, len(a))
// }

// func bisectLeftRange(a []int, v int, lo, hi int) int {
// 	s := a[lo:hi]
// 	return sort.Search(len(s), func(i int) bool {
// 		return s[i] >= v
// 	})
// }

// Binary search for ints example
// https://play.golang.org/p/G1CH7kPojZ6
