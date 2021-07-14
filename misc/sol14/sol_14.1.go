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

func Sol14_1_v1(head1, head2 *adt.SLLNode) *adt.SLLNode {
	carryForw := 0
	data1 := 0
	data2 := 0
	var resHead *adt.SLLNode
	var prevResNode *adt.SLLNode
	for head1 != nil || head2 != nil {
		if head1 != nil {
			data1 = head1.Data.(int)
		} else {
			data1 = 0
		}
		if head2 != nil {
			data2 = head2.Data.(int)
		} else {
			data2 = 0
		}
		sum := carryForw + data1 + data2

		if sum >= 10 {
			carryForw = 1
		} else {
			carryForw = 0
		}

		sum = sum % 10

		tmpNode := adt.NewSLLNode(sum)

		if prevResNode == nil {
			resHead = tmpNode
			prevResNode = resHead
		} else {
			prevResNode.Next = tmpNode
			prevResNode = prevResNode.Next
		}

		if head1 != nil {
			head1 = head1.Next
		}
		if head2 != nil {
			head2 = head2.Next
		}
	}
	if carryForw > 0 {
		prevResNode.Next = adt.NewSLLNode(carryForw)
	}
	return resHead
}
