/*
sol_12_test.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package misc
package misc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol_12_InsertInSortedOrder(t *testing.T) {

	tcs := []struct {
		given []int
		want  List
	}{
		{[]int{3, 0, 4, 2, 0, 1}, List{0, 0, 1, 2, 3, 4}},
		{[]int{0, 0, 1, 2, 3, 4}, List{0, 0, 1, 2, 3, 4}},
	}

	for _, tc := range tcs {
		var l *List // Or, l := NewList()
		for _, val := range tc.given {
			l = l.InsertInSortedOrder(val)
		}
		assert.Equal(t, *l, tc.want)
	}
}

func TestSol_12(t *testing.T) {
	testcases := []struct {
		list      []int     // positive integers
		calMedian []bool    // whether to calculate median for that integer or not
		median    []float32 // expected median value at a particular time/index
	}{
		{[]int{}, []bool{}, []float32{}},
		{[]int{0}, []bool{true}, []float32{0.0}},
		{[]int{3, 0}, []bool{true, false}, []float32{3.0, 1.5}},
		{
			[]int{3, 0, 4, 2, 0, 1},
			[]bool{true, false, true, false, true, true},
			[]float32{3.0, 1.5, 3.0, 2.5, 2.0, 1.5},
		},
	}

	for _, tc := range testcases {
		var c *Context = NewContext()
		for i := range tc.list {
			m := Sol_12(c, tc.list[i], tc.calMedian[i])
			if tc.calMedian[i] {
				assert.Equal(t, tc.median[i], m)
			} else {
				assert.Equal(t, float32(-1.0), m)
			}
		}
	}
}
