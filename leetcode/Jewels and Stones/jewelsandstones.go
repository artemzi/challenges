package jewelsandstones

/**
You're given strings J representing the types of stones that are jewels, and S representing the stones you have.  Each character in S is a type of stone you have.  You want to know how many of the stones you have are also jewels.

	The letters in J are guaranteed distinct, and all characters in J and S are letters. Letters are case sensitive, so "a" is considered a different type of stone from "A".

	Example 1:

		Input: J = "aA", S = "aAAbbbb"
		Output: 3
	Example 2:

		Input: J = "z", S = "ZZ"
		Output: 0
	Note:

		* S and J will consist of letters and have length at most 50.
		* The characters in J are distinct.
*/

type MyHashMap struct {
	val map[rune]int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{
		make(map[rune]int),
	}
}

/** value will always be positive. */
func (m *MyHashMap) Put(key rune, value int) {
	m.val[key] = value
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (m *MyHashMap) Get(key rune) int {
	if val, ok := m.val[key]; ok {
		return val
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (m *MyHashMap) Remove(key rune) {
	delete(m.val, key)
}

func numJewelsInStones(J string, S string) int {
	var result int
	m := Constructor()
	for _, r := range J {
		m.Put(r, 0)
	}

	for _, r := range S {
		if ok := m.Get(r); ok != -1 {
			m.val[r]++
		}
	}

	for _, v := range m.val {
		result += v
	}
	return result
}
