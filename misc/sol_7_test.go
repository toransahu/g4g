/*
sol_7_test.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package misc
package misc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol7(t *testing.T) {
	testcases := []struct {
		given [][]int
		want  int
	}{
		{given: [][]int{
			{0, 1, 1, 1, 0, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 0, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 0, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 0, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
			{1, 1, 1, 1, 1, 0, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		}, want: 11},
	}

	for _, testcase := range testcases {
		assert.Equal(t, testcase.want, Sol7(testcase.given))
	}
}

func TestSol7_getCellPos(t *testing.T) {
	type Given struct {
		width int
		index int
	}
	type Want struct {
		row int
		col int
	}

	testcases := []struct {
		given Given
		want  Want
	}{
		{given: Given{width: 4, index: 3}, want: Want{row: 0, col: 3}},
		{given: Given{width: 4, index: 2}, want: Want{row: 0, col: 2}},
		{given: Given{width: 4, index: 5}, want: Want{row: 1, col: 1}},
		{given: Given{width: 4, index: 7}, want: Want{row: 1, col: 3}},
		{given: Given{width: 4, index: 0}, want: Want{row: 0, col: 0}},
	}

	for _, testcase := range testcases {
		row, col := getCellPos(testcase.given.index, testcase.given.width)
		assert.Equal(t, testcase.want.row, row, testcase)
		assert.Equal(t, testcase.want.col, col, testcase)
	}
}
