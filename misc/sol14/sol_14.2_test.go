/*
sol_14.2_test.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package sol14
package sol14

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/toransahu/goutils/adt"
)

func TestSol14_2(t *testing.T) {
	ll1 := adt.NewSLinkedList()
	ll1.Append(1)
	ll1.Append(2)
	ll1.Append(6)

	ll2 := adt.NewSLinkedList()
	ll2.Append(9)
	ll2.Append(8)
	ll2.Append(7)

	ll3 := adt.NewSLinkedList()
	ll3.Append(1)
	ll3.Append(1)
	ll3.Append(2)
	ll3.Append(6)
	ll3.Append(9)

	ll4 := adt.NewSLinkedList()
	ll4.Append(9)
	ll4.Append(8)
	ll4.Append(8)
	ll4.Append(7)

	testcases := []struct {
		given []adt.SLinkedlist
		want  []int
	}{
		{given: []adt.SLinkedlist{ll1, ll2}, want: []int{1, 1, 1, 3}},
		{given: []adt.SLinkedlist{ll3, ll4}, want: []int{2, 1, 1, 5, 6}},
	}

	for _, tc := range testcases {
		resHead := Sol14_2_v1(tc.given[0].Head, tc.given[1].Head)

		for _, val := range tc.want {
			assert.Equal(t, val, resHead.Data.(int))
			resHead = resHead.Next
		}
	}
}
