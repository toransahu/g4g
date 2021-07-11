/*
sol_13_test.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package sol13
package sol13

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol13(t *testing.T) {
	testcases := []struct {
		given int
		want  int
	}{
		{-1, 0},
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 4},
		{4, 7},
		{5, 13},
	}

	for _, tc := range testcases {
		got := Sol13(tc.given)
		assert.Equal(t, tc.want, got)
		got = Sol13_v2(tc.given)
		assert.Equal(t, tc.want, got)
		got = Sol13_v3(tc.given)
		assert.Equal(t, tc.want, got)
	}
}
