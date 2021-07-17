/*
sol24_test.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package sol24
package sol24

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol24(t *testing.T) {
	testcases := []struct {
		given [][]int
		want  [][]int
	}{
		{
			given: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}},
			want:  [][]int{{13, 9, 5, 1}, {14, 10, 6, 2}, {15, 11, 7, 3}, {16, 12, 8, 4}},
		},
		{
			given: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want:  [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			given: [][]int{{1, 2}, {3, 4}},
			want:  [][]int{{3, 1}, {4, 2}},
		},
	}

	for _, tc := range testcases {
		got := Sol24_v1(tc.given)
		assert.Equal(t, tc.want, got)
	}
}
