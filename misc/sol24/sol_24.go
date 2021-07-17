/*
sol_24.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package sol24
package sol24

// Time complexity: O(n^2)
// Space complexity: O(1)
func Sol24_v1(arr [][]int) [][]int {
	size := len(arr)

	var tmp, valAtNewCoord int
	for i := 0; i < size/2; i++ {
		for j := i; j < size-i-1; j++ {
			x := i
			y := j

			origX, origY := x, y
			isInitialRound := true
			for {
				if !isInitialRound && origX == x && origY == y {
					break
				}

				X, Y := newCoord(x, y, size)

				valAtNewCoord = arr[X][Y]
				if isInitialRound {
					arr[X][Y] = arr[x][y]
				} else {
					arr[X][Y] = tmp
				}
				tmp = valAtNewCoord
				x, y = X, Y
				isInitialRound = false
			}
		}
	}
	return arr
}

func newCoord(x, y, size int) (int, int) {
	return y, size - 1 - x
}
