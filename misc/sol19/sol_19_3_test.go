/*
sol_19_3_test.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package sol19
package sol19

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol19_3(t *testing.T) {
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

	tts := []struct {
		given *BSTNode
		want  bool
	}{
		{n0, true},
		{n10, false},
	}

	for _, tt := range tts {
		got := Sol19_3(tt.given)
		assert.Equal(t, tt.want, got)
	}
}
