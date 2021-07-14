/*
sol_14.2.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package sol14
package sol14

import (
	"github.com/toransahu/goutils/adt"
)

func Sol14_2_v1(head1, head2 *adt.SLLNode) *adt.SLLNode {
	carryForw := 0
	ctx := &Context{}

	// calculate the size of the ll1
	sizeLL1 := 0
	origHead1 := head1
	for head1 != nil {
		head1 = head1.Next
		ctx.sizeLL1++
	}
	head1 = origHead1

	// calculate the size of the ll2
	sizeLL2 := 0
	origHead2 := head2
	for head2 != nil {
		head2 = head2.Next
		ctx.sizeLL2++
	}
	head2 = origHead2

	// if ll2 is smaller than ll1
	if ctx.sizeLL1 > ctx.sizeLL2 {
		// append zeros to right of the number
		// so that both the numbers become of same number of digits
		head2 = addNDummyNodes(head2, ctx.sizeLL1-ctx.sizeLL2)

	} else if sizeLL1 < sizeLL2 {
		// append zeros to right of the number
		head1 = addNDummyNodes(head1, ctx.sizeLL2-ctx.sizeLL1)

	}

	// calculate the sum of the current 2 nodes/digits (right most / most significant ones)
	carryForw = sumNodes(ctx, head1, head2)

	// as the above was a recurrsive call, it has taken care of all the digits/places
	// so, after adding up most significant digits, if thereis any carry forward
	if carryForw > 0 {
		// then add that to right most of the resultant number/LL
		tmpNode := adt.NewSLLNode(carryForw)
		tmpNode.Next = ctx.prevResNode
		ctx.prevResNode = tmpNode
	}

	// return the head of the resultant LL
	return ctx.prevResNode
}

func addNDummyNodes(node *adt.SLLNode, n int) *adt.SLLNode {
	for n > 0 {
		tmpNode := adt.NewSLLNode(0)
		tmpNode.Next = node
		node = tmpNode
		n--
	}
	return node
}

func sumNodes(ctx *Context, node1, node2 *adt.SLLNode) int {
	if node1 == nil && node2 == nil {
		return 0
	}
	data1 := 0
	data2 := 0
	if node1 != nil {
		data1 = node1.Data.(int)
	}
	if node2 != nil {
		data2 = node2.Data.(int)
	}

	sum := data1 + data2 + sumNodes(ctx, node1.Next, node2.Next)

	carryForw := 0
	if sum >= 10 {
		carryForw = 1
	}
	sum = sum % 10

	tmpNode := adt.NewSLLNode(sum)

	if ctx.prevResNode == nil {
		ctx.prevResNode = tmpNode
	} else {
		tmpNode.Next = ctx.prevResNode
		ctx.prevResNode = tmpNode
	}

	return carryForw
}

type Context struct {
	prevResNode *adt.SLLNode
	sizeLL1     int
	sizeLL2     int
}
