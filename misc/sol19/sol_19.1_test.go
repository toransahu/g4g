/*
sol_19.1_test.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package sol19
package sol19

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol19_1(t *testing.T) {
	n6 := &BSTNode{data: 15}
	n5 := &BSTNode{data: 11}
	n4 := &BSTNode{data: 7}
	n3 := &BSTNode{data: 3}
	n2 := &BSTNode{data: 13, left: n5, right: n6}
	n1 := &BSTNode{data: 5, left: n3, right: n4}
	n0 := &BSTNode{data: 10, left: n1, right: n2}

	type Given struct {
		node *BSTNode
		k    int
	}

	tts := []struct {
		given Given
		want  int
	}{
		{Given{node: n0, k: 1}, 15},
		{Given{node: n0, k: 2}, 13},
		{Given{node: n0, k: 3}, 11},
		{Given{node: n0, k: 4}, 10},
	}

	for _, tt := range tts {
		got := Sol19_1(tt.given.node, tt.given.k)
		assert.Equal(t, tt.want, got)
	}
}
