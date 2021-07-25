/*
sol_19_test.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package sol19
package sol19

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol19(t *testing.T) {
	n6 := &BSTNode{data: 15}
	n5 := &BSTNode{data: 11}
	n4 := &BSTNode{data: 7}
	n3 := &BSTNode{data: 3}
	n2 := &BSTNode{data: 13, left: n5, right: n6}
	n1 := &BSTNode{data: 5, left: n3, right: n4}
	n0 := &BSTNode{data: 10, left: n1, right: n2}

	n16 := &BSTNode{data: 15}
	n15 := &BSTNode{data: 11}
	n14 := &BSTNode{data: 5}
	n13 := &BSTNode{data: 3}
	n12 := &BSTNode{data: 13, left: n15, right: n16}
	n11 := &BSTNode{data: 7, left: n13, right: n14}
	n10 := &BSTNode{data: 10, left: n11, right: n12}

	n26 := &BSTNode{data: 15}
	n25 := &BSTNode{data: 11}
	n24 := &BSTNode{data: 7}
	n23 := &BSTNode{data: 5}
	n22 := &BSTNode{data: 13, left: n25, right: n26}
	n21 := &BSTNode{data: 3, left: n23, right: n24}
	n20 := &BSTNode{data: 10, left: n21, right: n22}

	n36 := &BSTNode{data: 15}
	n35 := &BSTNode{data: 11}
	n34 := &BSTNode{data: 3}
	n33 := &BSTNode{data: 7}
	n32 := &BSTNode{data: 13, left: n35, right: n36}
	n31 := &BSTNode{data: 5, left: n33, right: n34}
	n30 := &BSTNode{data: 10, left: n31, right: n32}

	n46 := &BSTNode{data: 15}
	n45 := &BSTNode{data: 11}
	n44 := &BSTNode{data: 7}
	n43 := &BSTNode{data: 3}
	n42 := &BSTNode{data: 5, left: n45, right: n46}
	n41 := &BSTNode{data: 13, left: n43, right: n44}
	n40 := &BSTNode{data: 10, left: n41, right: n42}

	testcases := []struct {
		given *BSTNode
		want  []*BSTNode
	}{
		{n0, []*BSTNode{nil, nil}},
		{n10, []*BSTNode{n11, n14}},
		{n20, []*BSTNode{n23, n21}},
		{n30, []*BSTNode{n33, n34}},
		{n40, []*BSTNode{n41, n42}},
	}

	for _, tc := range testcases {
		first, second := Sol19v1(tc.given)
		assert.Same(t, tc.want[0], first)
		assert.Same(t, tc.want[1], second)

		first, second = Sol19v2(tc.given)
		assert.Same(t, tc.want[0], first)
		assert.Same(t, tc.want[1], second)
	}
}

func TestSol19_InOrderTraversal(t *testing.T) {

	n6 := &BSTNode{data: 15}
	n5 := &BSTNode{data: 11}
	n4 := &BSTNode{data: 7}
	n3 := &BSTNode{data: 3}
	n2 := &BSTNode{data: 13, left: n5, right: n6}
	n1 := &BSTNode{data: 5, left: n3, right: n4}
	bst1 := &BSTNode{data: 10, left: n1, right: n2}
	testcases := []struct {
		given *BSTNode
		want  []*BSTNode
	}{
		{bst1, []*BSTNode{n3, n1, n4, bst1, n5, n2, n6}},
	}

	for _, tc := range testcases {
		got := &[]*BSTNode{}
		InOrderTraversal(tc.given, got)
		assert.Equal(t, tc.want, *got)
	}
}
