package main

import (
	"slices"
)

func maxDistance(position []int, m int) int {
	slices.Sort(position)

	l := 1
	r := position[len(position)-1]
	ans := -1

	for l <= r {
		dist := (l + r) / 2
		if isValid(position, dist, m) {
			ans = dist
			l = dist + 1
		} else {
			r = dist - 1
		}
	}

	return ans
}

func isValid(position []int, dist int, m int) bool {
	prev := position[0]
	balls_placed := 1

	for _, pos := range position[1:] {
		if pos-prev >= dist {
			balls_placed++
			prev = pos
		}
	}

	return m <= balls_placed
}
