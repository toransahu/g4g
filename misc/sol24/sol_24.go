/*
sol_24.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package sol24
package sol24

// Sol24_v1 rotates the given arr by 90 degree (clockwise) by roating each circular/rectangular layer one by one.
// Time complexity: O(n^2)
// Space complexity: O(1)
func Sol24_v1(arr [][]int) [][]int {
	size := len(arr)

	for i := 0; i < size/2; i++ {
		for j := i; j < size-i-1; j++ {
			x := i
			y := j

			// top --> right
			nextX, nextY := nextCoord(x, y, size)
			valAtNewCoord := arr[nextX][nextY]
			arr[nextX][nextY] = arr[x][y]
			x, y = nextX, nextY
			tmp := valAtNewCoord

			// right --> bottom
			nextX, nextY = nextCoord(x, y, size)
			valAtNewCoord = arr[nextX][nextY]
			arr[nextX][nextY] = tmp
			x, y = nextX, nextY
			tmp = valAtNewCoord

			// bottom --> left
			nextX, nextY = nextCoord(x, y, size)
			valAtNewCoord = arr[nextX][nextY]
			arr[nextX][nextY] = tmp
			x, y = nextX, nextY
			tmp = valAtNewCoord

			// left --> top
			nextX, nextY = nextCoord(x, y, size)
			arr[nextX][nextY] = tmp
		}
	}
	return arr
}

func nextCoord(x, y, size int) (int, int) {
	return y, size - 1 - x
}
