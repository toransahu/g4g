/*
sol_14.1.go

Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package sol14
package sol14

import (
	"github.com/toransahu/goutils/adt"
)

// Sol14_1_v1 adds two numbers represented by two linked lists.
// Note: Least (right) significant digit is the head of the linkedlist.
// Approach: Itereative
// Time complexity: O(m or n)
// Space complexity: O(n) - to store the resultant LL
func Sol14_1_v1(head1, head2 *adt.SLLNode) *adt.SLLNode {
	carryForw := 0
	var resHead *adt.SLLNode
	var prevResNode *adt.SLLNode
	for head1 != nil || head2 != nil {
		// get the digit at current place in number 1
		data1 := 0
		if head1 != nil {
			data1 = head1.Data.(int)
		}

		// get the digit at current place in number 2
		data2 := 0
		if head2 != nil {
			data2 = head2.Data.(int)
		}

		// sum of two digits of two numbers = sum of current digits + carry forward if any
		sum := carryForw + data1 + data2

		// calculate next carry forward
		if sum >= 10 {
			carryForw = 1
		} else {
			carryForw = 0
		}

		// digit to be placed in resultant number (from current sum)
		sum = sum % 10

		// place the digit in the resultant number
		tmpNode := adt.NewSLLNode(sum)
		if prevResNode == nil {
			resHead = tmpNode
			prevResNode = resHead
		} else {
			prevResNode.Next = tmpNode
			prevResNode = prevResNode.Next
		}

		// point to next digit in the number-1
		if head1 != nil {
			head1 = head1.Next
		}

		// point to next digit in the number-2
		if head2 != nil {
			head2 = head2.Next
		}
	}

	// at the end, if any carry forward, place that digit to the resultant number
	if carryForw > 0 {
		prevResNode.Next = adt.NewSLLNode(carryForw)
	}

	return resHead
}
