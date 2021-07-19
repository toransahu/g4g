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

// Sol14_2_v1 adds two numbers represented by two linked lists.
// Note: Most (left) significant digit is the head of the linkedlist.
// Approach: Recursive
// Time complexity: O(m or n)
// Space complexity: O(n) - to store the resultant LL
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

// addNDummyNodes prefixes the LL (representing a number) by placing nodes with value=zero at head (left) side.
// Time complexity: O(n) [here n is given as a constant]
func addNDummyNodes(node *adt.SLLNode, n int) *adt.SLLNode {
	for n > 0 {
		tmpNode := adt.NewSLLNode(0)
		tmpNode.Next = node
		node = tmpNode
		n--
	}
	return node
}

// sumNodes adds two numbers represented by two linked lists and returns the carry forward.
// Note: Most (left) significant digit is the head of the linkedlist.
// Approach: recurrsive
// Time complexity: O(n)  [from T(n) = T(n-1) + O(1)]
// Space complexity: O(n)
func sumNodes(ctx *Context, node1, node2 *adt.SLLNode) int {
	// if nodes are null
	if node1 == nil && node2 == nil {
		// the sum is zero
		return 0
	}

	// assume digits are zero
	data1 := 0
	data2 := 0
	// if nodes have values, assign that
	if node1 != nil {
		data1 = node1.Data.(int)
	}
	if node2 != nil {
		data2 = node2.Data.(int)
	}

	// sum of 2 digits in a number = sum of current digits + carry-forward from previous addition
	sum := data1 + data2 + sumNodes(ctx, node1.Next, node2.Next)

	// calculate the carry forward
	carryForw := 0
	if sum >= 10 {
		carryForw = 1
	}

	// effective sum from addition of current digits
	sum = sum % 10

	// place the effective sum value in the resultant linkedlist
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
