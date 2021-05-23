/*
sol_8_test.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package misc
package misc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/toransahu/goutils/adt"
)

func TestSol8(t *testing.T) {
	/*
			[1] → [2] → [3]
		           ↑     ↓
		          [5] ← [4]
	*/
	l1 := adt.NewSLinkedList()
	l1.Append(1)
	node2 := l1.Append(2)
	l1.Append(3)
	l1.Append(4)
	node5 := l1.Append(5)
	node5.Next = node2

	/*
			[1] → [2] → [3]
		                 ↓
		          [5] ← [4]
	*/
	l2 := adt.NewSLinkedList()
	l2.Append(1)
	l2.Append(2)
	l2.Append(3)
	l2.Append(4)
	l2.Append(5)

	/*
		<empty/nil>
	*/
	l3 := adt.NewSLinkedList()

	/*
		[1]
	*/
	l4 := adt.NewSLinkedList()
	l4.Append(1)

	/*
		   [1]-->
		    ↑   ↓
			<----
	*/
	l5 := adt.NewSLinkedList()
	l5.Append(1)
	l5.Head.Next = l5.Head

	/*
			[1] → [2] → [3]
		     ↑           ↓
		    [6] ← [5] ← [4]
	*/
	l6 := adt.NewSLinkedList()
	l6.Append(1)
	l6.Append(2)
	l6.Append(3)
	l6.Append(4)
	l6.Append(5)
	node6 := l6.Append(6)
	node6.Next = l6.Head

	testcases := []struct {
		given  *adt.SLinkedlist
		wanted bool
	}{
		{given: &l1, wanted: true},
		{given: &l2, wanted: false},
		{&l3, false},
		{&l4, false},
		{&l5, true},
		{&l6, true},
	}

	for _, tc := range testcases {
		got := Sol8(tc.given)
		assert.Equal(t, tc.wanted, got)

		got, l := Sol8_v2(tc.given)
		assert.Equal(t, tc.wanted, got)
		for node := range l.Iterate() {
			fmt.Println(node.Data)
		}
		fmt.Println()
	}
}
