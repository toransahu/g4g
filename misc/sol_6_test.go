/*
sol_6_test.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package misc
package misc

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/toransahu/goutils/adt"
)

func TestSol6_v3(t *testing.T) {
	testcases := []struct {
		given  [][]int
		wanted []int
	}{
		{given: [][]int{{1, 3, 7}, {2, 4, 6}, {0, 5, 8}}, wanted: []int{0, 1, 2, 3, 4, 5, 6, 7, 8}},
		{given: [][]int{{1, 3, 7}}, wanted: []int{1, 3, 7}},
		{given: [][]int{{1, 3, 7}, {}, {}}, wanted: []int{1, 3, 7}},
		{given: [][]int{{1, 3, 7}, {0}, {}}, wanted: []int{0, 1, 3, 7}},
		{given: [][]int{{1, 3, 7}, {0}, {-1}}, wanted: []int{-1, 0, 1, 3, 7}},
		{given: [][]int{{1, 2, 7}, {3, 8, 10}, {0, 5, 9}}, wanted: []int{0, 1, 2, 3, 5, 7, 8, 9, 10}},
		{given: [][]int{{1, 2, 7}, {2, 8, 10}, {0, 7, 9}}, wanted: []int{0, 1, 2, 2, 7, 7, 8, 9, 10}},
	}

	for _, testcase := range testcases {
		got := Sol_6_v3(testcase.given)
		if !reflect.DeepEqual(got, testcase.wanted) {
			t.Errorf("given: %v, got: %v, want: %v", testcase.given, got, testcase.wanted)
		}
	}
}

func TestSol6_v1(t *testing.T) {
	testcases := []struct {
		given  [][]int
		wanted []int
	}{
		{given: [][]int{{1, 3, 7}, {2, 4, 6}, {0, 5, 8}}, wanted: []int{0, 1, 2, 3, 4, 5, 6, 7, 8}},
		{given: [][]int{{1, 3, 7}}, wanted: []int{1, 3, 7}},
		{given: [][]int{{1, 3, 7}, {}, {}}, wanted: []int{1, 3, 7}},
		{given: [][]int{{1, 3, 7}, {0}, {}}, wanted: []int{0, 1, 3, 7}},
		{given: [][]int{{1, 3, 7}, {0}, {-1}}, wanted: []int{-1, 0, 1, 3, 7}},
	}

	for _, testcase := range testcases {
		got := Sol_6_v1(testcase.given)
		if !reflect.DeepEqual(got, testcase.wanted) {
			t.Errorf("given: %v, got: %v, want: %v", testcase.given, got, testcase.wanted)
		}
	}
}

func TestInsertInSortedOrder(t *testing.T) {
	l := adt.NewSLinkedList()
	InsertInSortedOrder(&l, adt.NewSLLNode(1))
	InsertInSortedOrder(&l, adt.NewSLLNode(3))
	InsertInSortedOrder(&l, adt.NewSLLNode(7))
	InsertInSortedOrder(&l, adt.NewSLLNode(2))
	InsertInSortedOrder(&l, adt.NewSLLNode(4))
	InsertInSortedOrder(&l, adt.NewSLLNode(6))
	InsertInSortedOrder(&l, adt.NewSLLNode(0))
	InsertInSortedOrder(&l, adt.NewSLLNode(5))
	InsertInSortedOrder(&l, adt.NewSLLNode(8))

	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	got := []int{}
	for n := range l.Iterate() {
		got = append(got, n.Data.(int))
	}
	assert.Equal(t, want, got)
}
