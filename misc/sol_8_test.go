/*
sol_8_test.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package misc
package misc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/toransahu/goutils/adt"
)

func TestSol8(t *testing.T) {
	l := adt.NewSLinkedList()
	l.Append(1)
	node2 := l.Append(2)
	l.Append(3)
	l.Append(4)
	node5 := l.Append(5)
	node5.Next = node2

	/*
			[1] → [2] → [3]
		           ↑     ↓
		          [5] ← [4]
	*/

	assert.True(t, Sol8(&l))

	l = adt.NewSLinkedList()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Append(5)

	assert.False(t, Sol8(&l))
}
