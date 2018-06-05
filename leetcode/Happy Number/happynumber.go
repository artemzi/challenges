package happynumber

/**
Write an algorithm to determine if a number is "happy".

	A happy number is a number defined by the following process: Starting with any positive integer, replace the number by the sum of the squares of its digits, and repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1. Those numbers for which this process ends in 1 are happy numbers.

	Example:

	Input: 19

	Output: true

	Explanation:
		12 + 92 = 82
		82 + 22 = 68
		62 + 82 = 100
		12 + 02 + 02 = 1
*/

import (
	"strconv"
)

type MyHashSet struct {
	val []int
}

// Use code from previous task

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{make([]int, 0, 2^20)}
}

func (m *MyHashSet) Add(key int) {
	if !m.Contains(key) {
		m.val = append(m.val, key)
	}
}

func (m *MyHashSet) Remove(key int) {
	if k := m.Has(key); k != -1 {
		m.val = append(m.val[:k], m.val[k+1:]...)
	}
}

// Has is same as Contains but return element index
func (m *MyHashSet) Has(key int) int {
	for i, v := range m.val {
		if v == key {
			return i
		}
	}
	return -1
}

/** Returns true if this set did not already contain the specified element */
func (m *MyHashSet) Contains(key int) bool {
	for _, v := range m.val {
		if v == key {
			return true
		}
	}
	return false
}

// ======

func isHappy(n int) bool {
	s := &MyHashSet{}
	for n != 1 {
		if s.Contains(n) {
			return false
		}
		s.Add(n)
		n = calculate(n)
	}
	return true
}

func calculate(n int) int {
	tmp := strconv.Itoa(n)
	var res int

	for i := range tmp {
		val, _ := strconv.Atoi(string(tmp[i]))
		res += val * val
	}
	return res
}
