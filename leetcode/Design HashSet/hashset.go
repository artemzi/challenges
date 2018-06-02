/**
Design HashSet
	https://leetcode.com/explore/learn/card/hash-table/182/practical-applications/1139/

	Design a HashSet without using any built-in hash table libraries.

	To be specific, your design should include these two functions:

		add(value): Insert a value into the HashSet.
		contains(value) : Return whether the value exists in the HashSet or not.
		remove(value): Remove a value in the HashSet. If the value does not exist in the HashSet, do nothing.

	Example:

	MyHashSet hashSet = new MyHashSet();
		hashSet.add(1);
		hashSet.add(2);
		hashSet.contains(1);    // returns true
		hashSet.contains(3);    // returns false (not found)
		hashSet.add(2);
		hashSet.contains(2);    // returns true
		hashSet.remove(2);
		hashSet.contains(2);    // returns false (already removed)

	Note:

		All values will be in the range of [1, 1000000].
		The number of operations will be in the range of [1, 10000].
		Please do not use the built-in HashSet library.
*/
package hashset

type MyHashSet struct {
	val []int
}

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

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
