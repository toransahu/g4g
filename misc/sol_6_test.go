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

func TestSol6_v1(t *testing.T) {
	testcases := []struct {
		given  [][]int
		wanted []int
	}{
		{given: [][]int{{1, 3, 7}, {2, 4, 6}, {0, 5, 8}}, wanted: []int{0, 1, 2, 3, 4, 5, 6, 7, 8}},
	}

	for _, testcase := range testcases {
		got := Sol_6_v1(testcase.given)
		if !reflect.DeepEqual(got, testcase.wanted) {
			t.Errorf("given: %v, got: %v, want: %v", testcase.given, got, testcase.wanted)
		}
	}
}

func Test_insertInSortedOrder(t *testing.T) {
	l := adt.NewSLinkedList()
	insertInSortedOrder(&l, adt.NewSLLNode(1))
	insertInSortedOrder(&l, adt.NewSLLNode(3))
	insertInSortedOrder(&l, adt.NewSLLNode(7))
	insertInSortedOrder(&l, adt.NewSLLNode(2))
	insertInSortedOrder(&l, adt.NewSLLNode(4))
	insertInSortedOrder(&l, adt.NewSLLNode(6))
	insertInSortedOrder(&l, adt.NewSLLNode(0))
	insertInSortedOrder(&l, adt.NewSLLNode(5))
	insertInSortedOrder(&l, adt.NewSLLNode(8))

	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	got := []int{}
	for n := range l.Iterate() {
		got = append(got, n.Data.(int))
	}
	assert.Equal(t, want, got)
}
