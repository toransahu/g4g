/*
sol_11.1_test.go
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

func TestSol_11_1(t *testing.T) {

	l1 := adt.NewSLinkedList()
	node1 := l1.Append(3)
	node2 := l1.Append(0)
	node3 := l1.Append(4)
	node4 := l1.Append(2)
	node5 := l1.Append(10)
	node6 := l1.Append(1)
	node7 := l1.Append(11)
	l1.Append(0)
	node9 := l1.Append(9)
	node10 := l1.Append(7)
	node11 := l1.Append(14)

	resultHead := Sol11_1(&l1)
	resultLL := adt.NewSLinkedList()
	resultLL.Head = resultHead

	sl1 := adt.NewSLinkedList()
	sl1.Head = node2
	node2.Next = node6
	node6.Next = node4
	node4.Next = node1
	node1.Next = node3
	node3.Next = node10
	node10.Next = node9
	node9.Next = node5
	node5.Next = node7
	node7.Next = node11
	node11.Next = nil

	gotNode := resultLL.Head
	for wantNode := range sl1.Iterate() {
		fmt.Println(wantNode.Data.(int))
		assert.Equal(t, wantNode.Data.(int), gotNode.Data.(int))
		gotNode = gotNode.Next
	}
}
