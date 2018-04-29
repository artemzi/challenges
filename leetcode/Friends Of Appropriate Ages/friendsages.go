// Package friendsages solve contest problem
// https://leetcode.com/contest/weekly-contest-82/problems/friends-of-appropriate-ages/
package friendsages

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

func checkConditions(A int, B int) bool {
	if A < 100 && B > 100 {
		return false
	}

	if B > A || float64(B) <= 0.5*float64(A)+7 {
		return false
	}

	return true
}

func numFriendRequests(ages []int) int {
	// TODO: fix ' Time Limit Exceeded' on leetcode
	var requests int

	for i := 0; i < len(ages); i++ {
		for j := i + 1; j < len(ages); j++ {
			if checkConditions(ages[i], ages[j]) {
				requests++
			}
		}
	}

	for i := len(ages) - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if checkConditions(ages[i], ages[j]) {
				requests++
			}
		}
	}

	return requests
}
