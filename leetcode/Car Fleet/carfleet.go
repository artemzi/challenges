package main

import (
	"fmt"
	"log"
	"sort"
)

/**
N cars are going to the same destination along a one lane road.
	The destination is target miles away.

	Each car i has a constant speed speed[i] (in miles per hour), and initial position
	position[i] miles towards the target along the road.

	A car can never pass another car ahead of it, but it can catch up to it, and drive
	bumper to bumper at the same speed.

	The distance between these two cars is ignored - they are assumed to have the same
	position.

	A car fleet is some non-empty set of cars driving at the same position and same speed.
	Note that a single car is also a car fleet.

	If a car catches up to a car fleet right at the destination point, it will still
	be considered as one car fleet.


	How many car fleets will arrive at the destination?

	Note:

		0 <= N <= 10 ^ 4
		0 < target <= 10 ^ 6
		0 < speed[i] <= 10 ^ 6
		0 <= position[i] < target
		All initial positions are different.
*/

type Car struct {
	position int
	time     float64
}

type By func(c1, c2 *Car) bool

type carSorter struct {
	cars []*Car
	by   func(c1, c2 *Car) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (s *carSorter) Len() int {
	return len(s.cars)
}

// Swap is part of sort.Interface.
func (s *carSorter) Swap(i, j int) {
	s.cars[i], s.cars[j] = s.cars[j], s.cars[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *carSorter) Less(i, j int) bool {
	return s.by(s.cars[i], s.cars[j])
}

func (by By) Sort(cars []*Car) {
	cr := &carSorter{
		cars: cars,
		by:   by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(cr)
}

func carFleet(target int, position []int, speed []int) int {
	var ans int
	N := len(position)
	cars := make([]*Car, N)
	for i, pos := range position {
		cars[i] = &Car{
			pos,
			float64(target-pos) / float64(speed[i]),
		}
	}

	By(func(c1, c2 *Car) bool {
		return c1.position < c2.position
	}).Sort(cars)

	for _, c := range cars {

		log.Printf("%v:%v\n", c.position, c.time)
	}

	t := N - 1
	for t > 0 {
		if cars[t].time < cars[t-1].time {
			ans++
		} else {
			cars[t-1] = cars[t]
		}
		t--
	}

	if 0 == t {
		return ans + 1
	}

	return ans
}

func main() {
	fmt.Printf("expected: %v\ngot: %v\n", 3, carFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}))
	fmt.Printf("expected: %v\ngot: %v\n", 2, carFleet(10, []int{6, 8}, []int{3, 2}))
}
